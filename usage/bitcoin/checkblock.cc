#include <iostream>
#include <fstream>
#include <iomanip>
#include <endian.h>
#include <stdint.h>
#include <string.h>


const char block_start_message[4] = {0x0b, 0x11, 0x09, 0x07};
const int buffer_length = 1024 * 1024 * 256;

int getpos(const char* buffer, int length, const char* dest, int exlen) {
	int counter = 0;
	if(length <= 0) {
		return -1;
	}
	for(int i = 0;i < length;i ++) {
		if(buffer[i] == dest[counter]) {
			counter ++;
		} else {
			counter = 0;
		}

		if(counter == exlen) {
			return (i + 1) - exlen;
		}
	}

	return -1;
}

uint16_t getint16(const char* buffer) {
	union {
		uint16_t u16;
		uint8_t arr[2];
	} x;

	for(int i = 0;i < 2;i ++) {
		x.arr[i] = buffer[i];
	}
	uint16_t nsize = le16toh(x.u16);
	return nsize;
}

uint32_t getint32(const char* buffer) {
	union {
		uint32_t u32;
		uint8_t arr[4];
	} x;

	for(int i = 0;i < 4;i ++) {
		x.arr[i] = buffer[i];
	}
	uint32_t nsize = le32toh(x.u32);
	return nsize;
}

uint64_t getint64(const char* buffer) {
	union {
		uint64_t u64;
		uint8_t arr[8];
	} x;

	for(int i = 0;i < 8;i ++) {
		x.arr[i] = buffer[i];
	}
	uint64_t nsize = le64toh(x.u64);
	return nsize;
}

size_t getsizet(const char* buffer) {
	if(sizeof(size_t) == 4) {
		return getint32(buffer);
	} else {
		return getint64(buffer);
	}
}

uint64_t getvsize(const char* buffer, int& length) {
	uint8_t size_ind = buffer[0];
	uint64_t size = 0;
	if(size_ind < 253) {
		length = 1;
		size = size_ind;
	} else if(size_ind == 253) {
		length = 3;
		size = getint16(buffer + 1);
	} else if(size_ind == 254) {
		length = 5;
		size = getint32(buffer + 1);
	} else if(size_ind == 255) {
		length = 9;
		size = getint64(buffer + 1);
	}
	return size;
}


int sprintf_hex_r(const char* buffer, int length, char* dest) {
	dest += sprintf(dest, "0x");
	for(int i = 0;i < length;i ++) {
		dest += sprintf(dest, "%02x", uint8_t(buffer[length - i - 1]));
	}
}

int sprintf_hex(const char* buffer, int length, char* dest) {
	dest += sprintf(dest, "0x");
	for(int i = 0;i < length;i ++) {
		dest += sprintf(dest, "%02x", uint8_t(buffer[i]));
	}
}


int main() {

	//
	std::ifstream in;
	in.open("blk00001.dat", std::ifstream::in | std::ifstream::binary);
	if(!in.is_open()) {
		printf("Error opening file\n");
		return 1;
	}

	in.seekg(0, in.end);
	int in_size = in.tellg();
	printf("block file size : %d\n", in_size);

	//
	in.seekg(0, in.beg);
	char* buffer = new char[buffer_length];
	while(!in.eof()) {
		in.read(buffer, buffer_length);

		int buffer_pos = 0;
		int counter = 0;
		int bignore = 0;
		do {
			//
			int offset = getpos(buffer + buffer_pos, buffer_length - buffer_pos, block_start_message, 4);
			if(offset == -1) {
				printf("there is no block\n");
				break;
			}
			buffer_pos += offset;
			std::cout << "pos of block in buffer is : " << buffer_pos << std::endl;

			//
			const char* block = buffer + buffer_pos;
			int block_pos = 0;
			printf("block prefix : 0x%02x%02x%02x%02x\n", block[block_pos], block[block_pos+1], block[block_pos+2], block[block_pos+3]);
			block_pos += 4;

			//
			int nsize = getint32(block + block_pos);
			printf("block size : %d\n", nsize);
			block_pos += 4;

			//
			/*
			char* blockdata = new char[1024 * 1024 * 8];
			sprintf_hex_r(block, 8 + nsize, blockdata);
			printf("%s\n", blockdata);
			delete [] blockdata;
			*/

			//
			int version = getint32(block + block_pos);
			block_pos += 4;
			char hashPrevBlock[256];
			sprintf_hex_r(block + block_pos, 32, hashPrevBlock);
			block_pos += 32;
			char hashMerkleRoot[256];
			sprintf_hex_r(block + block_pos, 32, hashMerkleRoot);
			block_pos += 32;
			int nTime = getint32(block + block_pos);
			block_pos += 4;
			int nBits = getint32(block + block_pos);
			block_pos += 4;
			uint32_t nNonce = getint32(block + block_pos);
			block_pos += 4;
			printf("block header { version : %d, hashPrevBlock : %s, hashMerkleRoot : %s, nTime : %d, nBits : %08x, nNonce : %u}\n",
					version, hashPrevBlock, hashMerkleRoot, nTime, nBits, nNonce);

			//
			int vlength = 0;
			uint64_t tr_number = getvsize(block + block_pos, vlength);
			block_pos += vlength;
			printf("transation number in block is : %d\n", tr_number);

			//
			for(int i = 0;i < tr_number;i ++) {
				printf("xxxxxxxxxxxx transactionindex : %d\n",i);

				uint32_t txVersion = getint32(block + block_pos);
				printf("transaction version : %d\n", txVersion);
				block_pos += 4;

				int in_length = 0;
				uint64_t in_number = getvsize(block + block_pos, in_length);
				printf("transaction in size : %d\n", in_number);
				block_pos += in_length;

				for(int j = 0;j < in_number;j ++) {
					printf("xxxxxxxxxxxx transactioninindex : %d\n", j);
					//
					char hashOutput[256];
					sprintf_hex_r(block + block_pos, 32, hashOutput);
					printf("output hash : %s\n", hashOutput);
					block_pos += 32;
					uint32_t nOutput = getint32(block + block_pos);
					printf("output n : %u\n", nOutput);
					block_pos += 4;
					int sign_length = 0;
					uint64_t sign_size = getvsize(block + block_pos, sign_length);
					printf("signature script size : %d\n", sign_size);
					block_pos += sign_length;
					//char* signature_script = new char[1024 * 1024];
					//sprintf_hex_r(block + block_pos, sign_size, signature_script);
					//printf("signature script : %s\n", signature_script);
					block_pos += sign_size;
					//delete [] signature_script;
					uint32_t nSequence = getint32(block + block_pos);
					printf("sequence : %u\n", nSequence);
					block_pos += 4;
				}

				int out_length = 0;
				uint64_t out_number = getvsize(block + block_pos, out_length);
				printf("transaction out size : %d\n", out_number);
				block_pos += out_length;

				for(int j = 0;j < out_number;j ++) {
					printf("xxxxxxxxxxxx transactionoutindex : %d\n", j);

					uint64_t nAmount = getint64(block + block_pos);
					printf("amount : %lu\n", nAmount);
					block_pos += 8;
					int sign_length = 0;
					uint64_t sign_size = getvsize(block + block_pos, sign_length);
					printf("script size : %d\n", sign_size);
					block_pos += sign_length;
					//char* signature_script = new char[1024 * 1024];
					//sprintf_hex_r(block + block_pos, sign_size, signature_script);
					//printf("script : %s\n", signature_script);
					block_pos += sign_size;
					//delete [] signature_script;
				}

				uint32_t nLockTime = getint32(block + block_pos);
				printf("transaction LockTime : %d\n", nLockTime);
				block_pos += 4;
			}

			/*
			if(counter == 752 && bignore < 781 - 753) {
				bignore ++;
				printf("xxxxxxxxxxxx block counter : %d\n", counter);
			} else {
				counter ++;
				printf("xxxxxxxxxxxx block counter : %d\n", counter);
			}

			if(counter == 2000) {
				break;
			}
			*/

			counter ++;
			printf("xxxxxxxxxxxx block counter : %d\n", counter);

			//
			buffer_pos = buffer_pos + 8 + nsize;

		} while(true);

		return 0;
	}

	//
	in.close();
}
