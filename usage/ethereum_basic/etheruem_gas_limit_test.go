package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"testing"
)

func TestGetTransactionGasLimit(t *testing.T) {
	txlists := `ethereum	10000LADYS	0xcd0143567fa448763f22db1c526544dd4efda7c1695499ec6d89606e2f16263e
ethereum	1000VINU	0x1a80ac999950f4a5c1c3baaaadcac52068bbb49a9ddd75d8d5ee421654de44fb
ethereum	1000VOLT	0xdc8bf4091f546cd3810a7249d95818a1cab766700338d57a60d1e61d60cf63c1
ethereum	1EARTH	0xba430fb468d931eaf802555dff8c564649094b3c3697bf5677fe550550769155
ethereum	1INCH	0x68ee8272673e4dd54e6c26961f343b7fd5086712495bf427c88bec6a05c09ea0
ethereum	AAVE	0x4dc406e6d4a1cb01e43098e6636d23f4a5b035d123749907871830a3eab6e54e
ethereum	ACH	0xadbe182f08292ff903c7f4bc182d3530fba833920c0eeae0d3c1a91f88f8d98f
ethereum	AGIX	0x7b47ef64e061bf0f02c2c7ebbeebc110ec4f8ac6f31b58e55c83cf0a4c48dcc6
ethereum	AGLD	0x32a8d95f7491a97da2e464aafbb3810e18ba9080b72bc17182057d082433f027
ethereum	ALICE	0xf6a9cdaf7ccbec1dfebbaa8ff73ec788644803d0a1f75403f638b4d99172d258
ethereum	ALPHA	0xda90a427f821c03661c8ba5453432bfa394aa4210d8a13c40a63009200b4f01b
ethereum	AMP	0x0c0eaab147e7944fd7ac2ea976db2ff8be5ab45032727f522b0912ad994bb484
ethereum	ANKR	0x2e2cfcba3760b8d77eb9ce831d5c5ed23cfd6ea599c89b591ed9a9521e823b06
ethereum	ANT	0xa2f9c87b4ad080604b6b41d8cb4aec5f83b5f69be6dd715a60d450af5fa2c993
ethereum	APE	0xce218ed7918bcdab22c81dbca24fae7285af551347fd2f17cfd69f4b8c56668b
ethereum	API3	0xbfa0bf6f85f8560902d47f67539fea2f0dcf89d380b9eebcb670f85ac9d61ca5
ethereum	ARPA	0xd4c6ee9b72d3652f619f6fb7cf145f364c143cd8c86a75b9660f1d6934a470b9
ethereum	AUCTION	0xde622eb4fc1feb22af2c065ea6d23cd963d5d268cdf4985e601060c51abe628c
ethereum	AUDIO	0x560a7b89b09cb3da06090a2c55eaee12ccb968cca7120ac5aa6c40c75cd0281b
ethereum	AXS	0x382ded4b3349c04b8b45830b6e9c14f4c8f7aca21ff01a911217b94dd3ae996b
ethereum	BAL	0x57e2795617eca80ec9fbc667a2cb0192826c85abe07f500de8d6475adee73b8a
ethereum	BAND	0xe2563ea2094400399caa709b8f2de1380a6f64d5b12b6d7d101e7b584d37a5b2
ethereum	BAT	0x5cd5d9966880c9ecdfdc2be3b6a448a31747a646123a851e0790603aa8d78cd0
ethereum	BBF	0xb0ae93213a4f4a539e3c8171f0cf8b56385bcd521e13b68503bddd0d85f2e250
ethereum	BEND	0x5c57088e98bbc6d3e69aaf331d74ecd44067f37ec16c8600ee035ab8ab71dd9e
ethereum	BETA	0xf251c8392614829a387a6f70b07651a0c39c23642f47c54afc69556f35736ef2
ethereum	BICO	0xb13a64868c0cbafb4f24a961022db14562281be0b4ad6be79000d9c8d9be5b78
ethereum	BIT	0x4806e5f8f39be8002dcfb1fe54623f378593c7e7200835fbf4f50480df5e6335
ethereum	BLUR	0x945e0fbfdb1b9559641065f69a4e8fbe6036218655f7c8c93c4ff84e727c50e1
ethereum	BNT	0x164ef8006d1c043f96dfbbb468a42888721a08b92cb8976a3186db14078f418a
ethereum	BOND	0x718b50681eb48fc55858a988b0a3a8b462d4915d581a8ba2fc8c9782be7a59c3
ethereum	BPTC	0x666a1bd24e2592e8098a294a7a073d773bc31875ee2d13f4f27b67539376c594
ethereum	BRZ	0x85c53e1723c874c8e5f2bac8b7ff46ff2b51bb7825cf2370d18ea224fa43ad6e
ethereum	BUSD	0x516dbd8d888115653c79e712de89488445d2df3478225266e1f9b6a77f5ec4a2
ethereum	C98	0x0c0d83b1b0182fb167f8a39d3fd8e4df1062efc2ca9b9ff26e959b8b755e1d46
ethereum	CCX	0x03718d264342d0b3e3cf20ac6c2585da186962053634f2558e87c2baf44b9bea
ethereum	CEEK	0x0f242941cc922a7af3bb756adb1eb33d1a1f4734ac8192ea59cdcc13a90fc34b
ethereum	CEL	0x4670b2bf3bd026fcfe64dadf8ca49d2e2edd7b24ce57319e720cead874b0e8ba
ethereum	CELR	0xf1e6fcb748e918a43bfce14577f185ded36edc9a27e154fd7ef6420402c532ee
ethereum	CHR	0xa9c7b958f9a0d541eb759e50e3d9ed8fd9f3a0ba16ac44e374fc8b751578ca9c
ethereum	CHZ	0x8e66a45ea1284badbf69bfc3797aec2ee9b3ba8eb57e152a8efc8a2b35e68abb
ethereum	CLV	0x28f0ffe9df3a21a4d82ea8a16997c6b8b0c8905e4c201c51d90f7a815610df79
ethereum	COCOS	0xa9b83be2522abd846309af3c4445d81dc498bd70c3e7e17f5d098c3e9f50c817
ethereum	COMBO	0xad879b84a01a5d8476c387c6f78ee5cbbfaf237a3671f3f1fd4ded8704f8860b
ethereum	COMP	0xded98449001dfb45a0fcad937cece17899667a625f678ced1957482501dfbe74
ethereum	COTI	0x5d394467e46e431a02d02c14d61c128761b67168e6e4637cdbe1fae23ff8124b
ethereum	CRO	0xa25a5b6a2590c4250ed552e252646592e54ca29112ca3d7a2e51f9833b9d73a9
ethereum	CRPT	0x41c56d66ade59e44585278a47d731a601205e369526b6d90e29a9bc44a684e7a
ethereum	CRTS	0xe251352459754d70d3718dbd2901e9f464fb9647ec90fb536518c3713f68f9d0
ethereum	CRV	0x486666bb35a01c1e0efe713945c1eb84220796adc2b0b46b898d873e94faa68f
ethereum	CTSI	0x58bf1e250ae7cdd00d61f793cb4b96cf46f0b67e57aa072a73845e133aa79a82
ethereum	CVC	0x34151058ef1b600aae4b5a5623ed24bafab4f1781ea76910cf36901aca4a1d39
ethereum	CVP	0x2e5c74280918b4ee3cdcbb22edc72dd2118df729db812a5f6328b412a5899434
ethereum	CVX	0x00982ba5ad0b5da65ac15c94903bce9961b93c8f99ebbb6b0abb5530b88a5d50
ethereum	DAI	0x448f3f2cf431691347ecdf198d6a20cca530e636e3e686d1f961e5dd5b739be2
ethereum	DAO	0x1e1be9c26d6beab907a9899809a02bf265ad9e1d70d895992c53bfcb3cad8316
ethereum	DC	0x25266c16eedf81e3b2dd7a308bed48ab57f6000427396be3a909b8cb29aff4d5
ethereum	DENT	0x7b0729238bc6b213e6500158de53442dc4c45f052a2a89f420dea074cad0b5ed
ethereum	DF	0x683a21b58a2234e7b8425de76da5b542022e54388f36d03dec259f7915c43523
ethereum	DNT	0xa9a92147e8ebb804d3c646c9eed094e0cd397bf08a86cdd4d21d6165a014d883
ethereum	DPX	0xd8fcaa6cf3d4cb328c7553b3c6c38a43d0f30384049367863b1fcf015208e2b3
ethereum	DREP	0xb57756807b8ad0cccf348a64917be305ffb9af9801b1bdc20ce31990b0cdbebe
ethereum	DUSK	0x9e205c5d64a9d074fc56c739132dd7f91242a0883fe9081908ddc7767fa91ba5
ethereum	DYDX	0x36d4c8679d2077e067a58896fb6d9df254ed786a434c5ea3655c5a03e3ac1624
ethereum	ELF	0x6b898b91bc7100221b435e527f39f5e229e1a528d9bc6ab8444a31f1e50b97ae
ethereum	ENJ	0x4f0cb60c0eb918a4ace6bd92f2b3d3f1f5af86b9d5830f7289d953ac1f31949c
ethereum	ENS	0x11730406524d55b63fe93ecedc44ba8b423198ae3b24841d40b00afce7759961
ethereum	ERN	0x3fbcc8395de40524cf8537c8a0f14062ad42f1c3e33b7d8167faa2b4b7a0e130
ethereum	ETH	0xc10a0dbf7ac3f2870bdea1cc9a2924f8723806018d058779592b5e46ad2d7404
ethereum	ETHA	0xa8f4c1294c283e962ebe859ebc72ee7e47c6cc2904e9ccfc5c3cc2bb7ef1099d
ethereum	FARM	0x921946f320ff0ce66172b965d445d23b2f1a8dc063521c1b5d992c3980eb99ba
ethereum	FEI	0x3a637feebf85b3ebad93228a46c820deda0a58bc69f43bbd3de8ad38166b3bf9
ethereum	FET	0x089b4783b1cc5572bf071b3bf2b0c8b8ded9fa6b71c186d32e817b19090ad421
ethereum	FLUX	0xf625ec38ac931844ae7cf719fdce5988db064af0b38fdb0e3bc00f70fec066a1
ethereum	FOR	0xb44589f7815d61fac1b4b49d42990855e2121e1f9bca38f268f5ac7c1b9e3d03
ethereum	FTM	0x2c148c663ce5373990acc0a5412829806e99f2c105691c3f6f060afb650454e0
ethereum	FXS	0xa93554988eaf51a037b027ce5e1f37cc8946e38ce7ea0a8b54295bd7dd1a5a72
ethereum	GAL	0x1b624962622d5fc088d0471498eb478d942dd2a31d898336933585d1488d325e
ethereum	GALA	0xd4d1ef1a64ea8de17402cb8ccc8315e06d099b3f23903a33ec2d006956d5af8f
ethereum	GHST	0xe817843d08130ab6a704929bb883d99262e1a03ac2fee31f62bfc2fb5bb6487b
ethereum	GMTT	0x29a8da4b3a980266f5177ed0831adb15dc049d1bf490787c3a6d7cb4e557d103
ethereum	GNO	0x76aee17f4a89a0aa75634df728225b0f7e62996ca13b002db2ace05e422d0389
ethereum	GODS	0xf81294b5d7e79ac8502a3f125ae321c5c93cf18e9ab597ee9a84a7afb5c6bd38
ethereum	GRT	0xf88200ad17867cb82daa7248f8d2b9fce442281b6a2739809f8c161341fef719
ethereum	GTC	0x8a3a938c281c2d69276725cc7f48f25eb0cac9b1c0af5b22702b1dac8eac671a
ethereum	GTO	0x8fe998620da5913eaad5ff89548ac347eff5e9f93f044ffbfc85dbfc83930da9
ethereum	HFT	0xa32ecf714cbe32ee5d0026892fca56d88638b3a284ef97bebc456de91767e760
ethereum	HIGH	0x900c5883c69cb247232dd792e3809258ecbc1c7b9bba08e4d7689a50bac4cc9a
ethereum	HOT	0xc35e3e4291b5944c239196bda09f8d0f471c2d9142091c98a324189ce9341f63
ethereum	HT	0x4d122c457d6c4b705210769c88cd824930fbf1d0480643986d915350bed72360
ethereum	ID	0x2e7dd439f2565593e6a18a874401de91b843e84be9d103a6b711be0897d14d53
ethereum	IDEX	0x9f22da7973d271088d7fd7a7bf7154771b6e5d8bff81e8b2772cfcbe822cd30c
ethereum	ILV	0x6918ec7b93ca07a2989836e8311d50ecd44c1c1b1da2050702260077e365aa7b
ethereum	IMX	0x5cfe0bc3ce543b260269b5ce5cf07881e14d603fb71bd6ab53a3b9955f8ae8d7
ethereum	INJ	0x4369e06cddf633014971370c2a0506971feca9e55ba7ee011a5b23532c211aed
ethereum	JASMY	0xeedb9b7a1d93f74a6c3e4d50d031540027fe23a5d8f068ae549eaa7b2c0adabc
ethereum	KEY	0xd880701151891719b0223f5ee539638830a0ff9bae0dd2468651eacfce00998e
ethereum	KNC	0x150249f66f930f9045f41111855d78922cf825dddbcab1c09a99a964a05f0897
ethereum	KONO	0x677e289f973f285a31790bf44d3c157d2a17217c5e8278355c4b07ff65f80ded
ethereum	LDO	0xebf45bdb577c4791ff7e9f4009d7c29be4c87bd1b019c8052d726578ba96a418
ethereum	LEVER	0x74523582957e0836b0a1f2ae372bfe6681e76decc00f54c148230b7554259a7a
ethereum	LINK	0x9497a7eca32b89f9de1e107856c541da33922c6e53d0b1294cd50520bb66bfb9
ethereum	LIT	0xb50c606c9eeada6ed7adf85c8ccdaaedce30992a8263da8658c5948378b3ccf4
ethereum	LMR	0x13eacc3d1d010e0dc9335c3ab9435070a43c3e48862ac908838914930ae5a7aa
ethereum	LOKA	0x39c0ca3c838b25bc501c0ba504e30b5a982d47be0ce5fb79c9b18c12b0dc75e5
ethereum	LOOKS	0x6cbec41eff8668e37db239bbeb7f8fd9e1a9b8f6e6eadfe1bbc00b06e8aee3c0
ethereum	LOOT	0x3ca2a52bb5fd4aef86a9021e5212c93bc572b5eb7aba0627a961fa1a3326e6cb
ethereum	LPT	0x23241ea951ce04f28877192bcfb9ba809a57d72830808426a606a1d2c053766a
ethereum	LQTY	0xebe3c4b95c82758432340b1e1e956c0e2605af55285140c3436adaaa96c37704
ethereum	LRC	0xbd1bac6807ddbe770ddb3fe2826790da8022657414f679e3996228fe098f3874
ethereum	MANA	0xbfbd32b0b46577d56fa3437df36b9136bc66df6d78dd9552728945c23e8ecb2c
ethereum	MASK	0xdb5cb9062ba9681052c816118ef279d789755a664842c357253099c30ddef22a
ethereum	MATIC	0xb55f350b9649e700b07e9a31ef666a38e3ce8f514be586347ce6ce7a3252978f
ethereum	MDT	0x39052784a06f274b0a652a5d1614697bdeb33475f2eb099a2ca16ddc30c6d609
ethereum	MELOS	0x8211813a349aa701ba3a58a0f9d3d5912df7c2753f2df63337479830bc1a2e8f
ethereum	METIS	0x6a8b8a61459fbf23f32d35aa06cdf6312e3b221706ef412cb68d43cadb33104d
ethereum	MFT	0x3a34348a0704545c896c74b60b9245749947cc165b64331618f786280eaf0765
ethereum	MKR	0xbd872d01be18d2d9348a00db5ed67f3ed3b67e2970318fe668bee11333d15846
ethereum	MTL	0x80194028f942fb10086edfa0720c5bef3c71fbc5cfe794d7e4dc8257a72c2834
ethereum	MXC	0xb9393816b296d5ea8cfaba5ce448f2c896222d9852d9e59e76da6e00ad10fb6d
ethereum	NCT	0x06d5f831fe580c449df5aad5dc61729c102c73cd4875fd58b32dc20637723721
ethereum	NEST	0x91740d7b25baffb09723389e2b5965dde1ec7baf5d848c0fc484a5a76e3e8ac3
ethereum	NEXO	0x8fc81df8e58a5ab30f3408cea5beb51c5fec6f10d1e2e12a335146aaa7ab4031
ethereum	NKN	0x61c3f11aa81a0f37d14d8d3970848ca82ceb09333331956bdeb726fcd5059120
ethereum	NU	0x505bbeadf26122cf621af1c7f66d32890630678f5b26e82dd7509d47b74374ab
ethereum	NYM	0x89dc0b227461e3e179c1394bdd5ffeae7aee7a80475520febafee63b0b7c2a7f
ethereum	OCEAN	0xcf10d051122681dbec21c1d1c0681f722d14d22bb315ebe7f720b254bede0e04
ethereum	OGN	0xa3bad12cd19280b614f7b015888ad171fc96080370902650120d7bef41e61202
ethereum	OM	0x97d55047c0e04cb4fcaf20b5a1abd8c6113019eb4a16cf9dc0188a057eafc71c
ethereum	OMG	0xada3ad0d556154e507627ce55b234b1005c5a6fa62f2925a6d2681979eebba9d
ethereum	ONSTON	0xb49d33b596114ccd086e5211d69aa1357afb3b27378b6501f69aa445310106a8
ethereum	OOKI	0x7daf0c6a39b15deb01f38af015ec4b5561ea9935a5b99ce216ddb3b7ee3187d0
ethereum	ORBS	0x29177287716ebfbb8a17b52d438aa6664f60047ca435d663c667af380b65822f
ethereum	ORN	0xdc66e5b87dd761cf10c5670fc61bb650e446249931f256d581635aaf51d8d781
ethereum	OXT	0xc729bece615d5dc79a02f7d32d8da197319b5966dab487f708102f346a566035
ethereum	PAXG	0x60e6aada2f81a73dfe18d76ad68aefecd7795c8f151c05665b29672ee1f19eac
ethereum	PEOPLE	0xd01663b07f514e21d026db10c21145cb89af42855bcb366ebc9298e0d94b482d
ethereum	PEPE	0xc315c2276eaedbd334afdcb9c7d92b704807fae5534c61ffdf38a9acb4ce0ad3
ethereum	PERL	0xd09c5d01e94c1500efee3f0955d0e2f537e1accf78715aeee994bd15cfe42963
ethereum	PERP	0xfc31c7966c2912a122eeb08139cc5e4d903316cc75da635dbb3003644e03ef77
ethereum	PHA	0x0a0c219f65983d8d4fb74beabcff684f24bcc77cf3a9ad32b8dc3ba01aa008b2
ethereum	PLA	0xe3e68ef43110835a20b98d7e51b1eb1252fffd99645b74c27f58048f210502f3
ethereum	POLC	0x670626ec67fc32bf247ff658890e9e84fafe32b2fc9ef45e80094012da2e3058
ethereum	POLS	0xf2d7e1314f6189142cea70745498de3b174a121d36476c5e3e1753bdae0f73c9
ethereum	POLY	0x3e6bd909429b657ede76defad75991abb11a81402013ab276d4d84eff664344d
ethereum	POND	0xb412f204e77db157fb0cf955f83bc859354e06842696ccc12796709a457719cc
ethereum	PRMX	0x3b251e9a4a96816a888bd12ab98c702450ccd815eb682bc5a8d5be1612438f5b
ethereum	PSTAKE	0x9f14786885daab81477e5964e6472d4e90fc9f6642eb49fd6ffa09097cb7348f
ethereum	PUNDIX	0x68fae0e8cecb8f0d3963b86f4d11e843e94f9aa8749b2001d80f19d1bf615970
ethereum	QNT	0x662d40a7ae7174839ea76a4897c46178c53bebd05842664d7b3743890ed1f6e2
ethereum	QRDO	0xaf27adefd34480b6428f290bf41857ec385dd451ec51ee149e0def0715c3d0c8
ethereum	RACA	0x29728d79d3f70b3ccd07858335ad4800802acdc6f083b8f449e9c5a3a7dd1ff7
ethereum	RAD	0xb9df531841a9a933288c9b9f02e7415102032bbd8f707e16a8aab85764f62994
ethereum	RAMP	0xd040149b0f0031e78d6cf821557089798282ca616aea4b3011d71c5bb893b1f5
ethereum	RANKER	0xa68a4f41f65b891a7aae07de7260121ade01e826e6c2edecd4445351d206a6f0
ethereum	RARE	0xef5fc60df060d2ac13412341b48794dc9ced03c922abcbb58f3bd3428ed2e11d
ethereum	REAP	0x8b7a9783c1654687c9de3dff2dbdb80b92717f538aa93aeb19519dd04abb6937
ethereum	RED	0xb68bc6a036b0af4f9c14ea41faca26b9d861d368b77fa214dd30a10e0cc8b342
ethereum	REEF	0x021e35c4d29ecd8917450cbdc017a642c044ad5d0c390dcba3089f1b0a7b197e
ethereum	REN	0x38222354714a459f73ca0f560380fb074a500099ab4b10a60c09a70e590666e1
ethereum	RGT	0xa367ed89150578ea0a463374ff034a9dbe6a66db23482b82d3aca59e1ed998a8
ethereum	RLC	0x08f8fee182e7445e55fe1d52709d80d0aab2d8ea023424921ee2f6d61d5f7a70
ethereum	RLY	0xe0846c566742710623b9aea34e8231e3ab1f82ce666c9b35c95bb92058aec699
ethereum	RNDR	0x85d7c4b1f2d2060e1eea436fadd8a094a2f5de3393170abf2b310537219ed7a4
ethereum	RPL	0x9ef31e461977346a7221bbc7b22ab42ad1f00eb617dabdb8c80b5374a196e990
ethereum	RSR	0x47b5d33d953f76a61710898087977a746de5e477a02dde7e40d104030d4347fd
ethereum	RSS3	0x8489e41c5906663e925407ed4319bc3745e5585232088cb5d0162ccaf8cfc11a
ethereum	SAND	0x6976768c2201514dc5c9a81284bd5dbc6fc09f4bab71421789e605d66981a974
ethereum	SHIB	0x34de7cd5c116b3d177dc2a58ff9dfcab77736a0ad2aa86b92ba18f05ae24e23c
ethereum	SLP	0x9afa5723533039c22b18e022ba4a5b2db53db155aeabeba1aaba7bc08b7017a7
ethereum	SNT	0x75b41de74449c4f3db8601c7f56583140a3e935d7de113cfa223e3b639160e81
ethereum	SNX	0xcd92728841925d1c1dd49cc4c289277bc34d7afc12326e6d9d1670296ceb3ee8
ethereum	SOS	0x8a70a1e7a07ec7b8c2d4da1fb05e02454f70dcd5ee95aa13cd088a9d9ae5781b
ethereum	SPA	0x1203d0194d05ac40d8b4ed4e658a195cf3ff00372e11355f7cce6ad4b004f250
ethereum	SPELL	0xca54b80fae4554215b922f7abd508c0de053f1137db4efb2b3c273626e326046
ethereum	SSV	0x7c25c146aaf2ea45f01484cda54981deb3843b025f2e93e0a723ee32de048dde
ethereum	STETH	0x861c9343494e2da107ede48125f5011e43941b464b535f6e6c7eba4ddfa49a37
ethereum	STG	0xcda1358abf6a87c5f99603acddbc9521050a347b81990f433f17dea7795078ea
ethereum	STMX	0xa76abb9c2f9fd0b39f7fcbebc0b8bdac4b94117d9081ac137b2d17bfea52bd3a
ethereum	STORJ	0xffe46af53f8cbaa9993129d0e8bfb739e7d13d8f55d3ae0feece479f992c5345
ethereum	STPT	0xe1448a7f3dd49e0d1f5b823f807b5f1771d6d3b673fc926183256191c784b38c
ethereum	STRONG	0xa3c22c657564105d9dc7dc73bb51f58f42414f469385727e78db7906f12565cd
ethereum	SUPER	0x765f877a1a5698db529eedffca95b5868454715da2f1305a90c8ea683924d31f
ethereum	SUSHI	0xab664ca96a52d059ac5ea9c2ef9057674b83564562559ea0a2c5ca53916a67db
ethereum	SXP	0x00326e85511ca4ab4308ca948682d434ae6f2bfaf80a2309c88fb4504e922a7a
ethereum	SYN	0x1a8602cd63a64611c27417239564bbfce76d3d59dbd5cdbba1e9ae6a6f2d52d1
ethereum	T	0xf2f367a8d91d18280c8d320960e7fa92cee17ac4bb80783d3ff3a34d6783ac5c
ethereum	TEL	0x3338a085d917e589adc34c1f859a1674a511c9406cdb7758b560bd8b82451615
ethereum	THE	0x4d32b85d57a33dad61a1f8b0310e813f7767d4add10dcbe35db711c65060c7b1
ethereum	TKB	0xbcff0aea8e49be3edb3bed316c186039263b0d996a8b8fbb1c965a2dd7eab7f1
ethereum	TLM	0xb64ccfc7157290a62e71c52c71c638a661d37d01095441666a1a85c3a53d7ef9
ethereum	TOMI	0x4bafaeff93826138da185c402ea1a1a5b6acea33bd4a0e06e49430d8e47770a4
ethereum	TOMO	0xcdcb0d98569410794f2afb4f2075859fe09afd85dc8cced7f4201c4eb73e75b8
ethereum	TORN	0x39e24c44d160292c75bfe836900b1b9bd0b25f1d9ab7e675e8e2e11fad03d76a
ethereum	TRIBE	0x85d4d4d2f72beec9db48b3cfcb7522887217989cf9ed9a1a6631eeecfb611a24
ethereum	TRU	0xc3b5efda3b9b98b38597deae297010cc9e20084957b75da8856801e24d03cd8b
ethereum	TRVL	0x81775275599ccefeda67a220e9b28e9f726605c22f43a1e49319702db9d6ccdb
ethereum	TURBO	0x62f1c9def4706fb95d82f1f3fb1513cdbba9b12f256ce19f2bcf5f4bcc6cd52a
ethereum	TUSD	0x378a4bcc5d40b29f7f53e5157a04cca9ecfb3ebabf1dee105aa1ce7631b749f6
ethereum	TVK	0x2c4c7058c2758665fc914ea460b15217545b755cfc0278b412a4fbb4a0017682
ethereum	UMA	0x11d50b5a725998b80021199812e04bd841bb53f61e0dcdeef8b1ffa88625df1c
ethereum	UMEE	0x23b677545308a8a183eebd12c7a2a53fed3151d74dbda4425415f36505f8ce6a
ethereum	UNI	0xfdec4425d9e8b57ad0cd2a930544984275769bb6bb64d5f56033cde54b3eb4da
ethereum	USDC	0x712ead571312e29e04e8ef3677fdcf67bb7bdb97fb9cd68c258ca45b6d30a774
ethereum	USDP	0xd1534861f22c51c174074ad6e21ca1c4721cfa939ab0545218092313e71f9ca6
ethereum	USDT	0x7bcce37889cee66a0c97c5f14a5444c4eb511aab0bbeddbfe46454a5dd58ffbd
ethereum	UTK	0xf9288c327409f1051cf465b4ef88354a2e56d6284de4e0437409ddc4be0d18de
ethereum	VGX	0xe61fe4727ea50fd577726d7c2c1c54e42122c5ea9e4a56e327e2e9540e154d8a
ethereum	VPAD	0x79e1da862d5391fbb54fdacdd279bce7c0a3220bd73500c78834462647acbcc4
ethereum	VR	0x0633487b96e336b046ca2b6729ac251d8dcbc602645bcfb5786c6d4a7bb03cf6
ethereum	VRA	0x2c61762c628bd954c03272dd09f7998e3a1bf4d0c5175733a301e5c81fa020c5
ethereum	WBTC	0xe47b8b3c6fa7370ef805d907840889d5270aa24b54ba8848eb5805ba409c7a7d
ethereum	WCI	0xecd7168f577019e02c006c5c4cb79cd0ceafabd7f8c5286ebabcd4e56b5779d9
ethereum	XEN	0xc072e047c3655019b8bf9905896415aeca0db5d0667dd407e9a9b7d8f5d675f2
ethereum	XYO	0xce8c070fbe748315600fca03d771b811ad82eba22a45176e2ff7500c14b17afa
ethereum	YFI	0x08bab0b415545356ccae81e06f4417bba99151e00cb0ff310813ccdbd498f20a
ethereum	YGG	0x531f7f5c635f03055726403534f7c802ea8daca05b07267d3fc196be70f75517
ethereum	ZENI	0x77a4ced6faadaca90169de437b1901d6203233d49908ddceb6e13d16262719ab
ethereum	ZRX	0x743c3532590774f0b1740364c945af4ebfd541524fe207f5de8a90a630e9cc8e`
	/*
			txlists := `smartchain	1000000PIT	0xee5863e920ce40552c3587708bbfad71455f320bc8bbb032a49d33bcdea34d2b
		smartchain	1000BRISE	0x14e5f361aa6384f7d43c1440f9330984eaa96231ad20bef71895d0ad5cb5d561
		smartchain	1000FLOKI	0x9e6cc0802c68ede355b78ce9ba8c39ac327323d80e535a5fcc1e7fb2cbfd647d
		smartchain	1000VINU	0xdf566b4a0d2c75be3650e667fb6d1f021a2f6feae7a0ae26feb951a81f5f7fd4
		smartchain	1000VOLT	0x47eb95bdf51c0061605dbbcae4532623fd1c06c0a9a221a71bfd7151520c9177
		smartchain	ACH	0x784c2510201c1c1a32b5b704d90e922f9c1c10ff5fa3987a316e59ea25504768
		smartchain	ALPACA	0x086fcfa3761e734988f75bf76355190c887fcf9013a9f128a70826a32045351b
		smartchain	ALPHA	0x193b8c69c49865e9c78ff9ed0ed38fb119b7fea93a2235436c19a1024d34624a
		smartchain	ALPINE	0x4dc91f7e6eba2d50f671db55d0d7cc8125fe32feac381f321a7ca9bf3a218018
		smartchain	ASR	0x60597ef2707075665f4effa52d3796c4c7e43f6739b1c891fb9a02111f4b0e00
		smartchain	ATA	0xa3a0b603381803ba1041ac533c5d5aee4c691829e510e8ebb7118c2a1486fc41
		smartchain	ATM	0xb729e3eb53ade16eab50fe17b5c5c93f2207e6afed2e299ff30622d7fb881997
		smartchain	BAKE	0x8ab9cc1da472a7629d899c934d9689b9ae9ecb0b2d22f5646055c062061ce319
		smartchain	BEL	0xe029978f4e20921d998d837c4280dfcbae4e75c4e990c14a359904f42a1b48d2
		smartchain	BNB	0x80d61150ed9aa238cb5fc381ea7086e9aff74ea09326821ac9a629c93a09861b
		smartchain	BNX	0x3712bb6040a0cf25aa36eb349bd0778bd472400afe36fe6952b4ed154248dadd
		smartchain	BSW	0x330ee9e3c692caf80fd8e116102ee62873010a59031e24d775672dc729b2d781
		smartchain	BTC	0x3fc87584afa33b501cb2c2d36b882146de039b80427d9d03c7242d25ce613eaa
		smartchain	BUSD	0xb0b0fae32dd15300cefe862721ea732759cdefa6bc047fb171b644b5431e159a
		smartchain	CAKE	0x20586115afba80890ce0500d26bf7b5526d9ddba1c24fb92873b1aa4561dbefa
		smartchain	CEEK	0xf7eb51871b881c7133c493ee78d1d540714b9872fe2a9ec4929d6ade9210ee28
		smartchain	CHESS	0xc13497d8155e5c0a8e17ec9c1e712c5c0dcb697dee6124ecec7726591d4ffced
		smartchain	CLV	0xaf0b809dbfe1eac8d674740565bc598bee70b64d5c6cbfa8467241c7142bc9d9
		smartchain	COMBO	0x2f0d688e0dc115248d4753f6e7480224e1dfe732a9955d97de507d18ef5b5dbc
		smartchain	CRTS	0x2a33ef13f02e6f4882ab72bcf2c6a1ba39a4e31f348351bce6fa7224ed2761c5
		smartchain	DAR	0x8273df9faa0d15f87e8b86428d1db6406ef2bc70cc4c1e72c63c05d94b0e3575
		smartchain	EDU	0x4c2ee4a719f6448af9efd06b6da32d63e49ac3d98f04ef867e3bf5db317f9817
		smartchain	EPS	0x26735e37a6397727a9143fcca44db6ccca065fb39d04e50a968e4193d486bf95
		smartchain	EPX	0x018e08ba7147ad957d08be7ecfd05a96b2c3291728721d1bcd99f8f2f9064bb9
		smartchain	ETH	0x0038b9f1a6aedf1cfaed6e10bd12f8e7432ed29a1822ebfcd837ec013f67fdc4
		smartchain	FND	0xf970b75c9f53724b2f7dee564b5a29f63d457483ce909022534a6600ce7f1830
		smartchain	FRONT	0xb56d6ba5edf39ce1ae3b7f520e5e548d51386bbb3b2c97d72c8775ecb29ecdef
		smartchain	GAL	0xf128e15446757541279ee008cf3832d87acd1ade5be9d871badcbfb4f3d25227
		smartchain	GMM	0x0461da21b223a9e62e184826ffd26ff27ab89ff48b5b3a6b530e527dc6f06a51
		smartchain	GMTT	0x003c4a60c4e3da975864009c57ae646b5801d8911695caed24881d328c1568d8
		smartchain	HFT	0x296de2245455d98ed3cfb8ae9c7f339ad00f4bae5bbcf1060a0aca08e6c824a2
		smartchain	HIGH	0x23870a9f8b2069a73b33b9d88798965ba79b79255075ffc8a46009dce852005c
		smartchain	HOOK	0x7697045af4ecf6d95c56dbabe88bdce24fcc4c180716a4526bb8d3662ede7ceb
		smartchain	ID	0x2b727a354e54a8eeb0c59b1a3007c2eaf9a04328c2828b67316b440f089139dc
		smartchain	ING	0xb227be97f1e65a8a73834f7e2435ed8a46b19c7784eeb323ca4d60448a0d8fb0
		smartchain	JOT	0x5546466b46ba6ea204e348e1c269b68d117a65ce6bf149465e7f1bf197b7677d
		smartchain	JUV	0x2d1b4058014f6d3679d2696af255df01090c8ab9be0e2cc30a4c7ae1af2e061b
		smartchain	KMON	0x98f36c5025e3963849c89531490a1545823cb540fed59495ad05025cd54e0c54
		smartchain	LAZIO	0x01f8710ef439e8fa4c2f1772910a435c01683b81e228b2d9b52db3f67c7d7dcb
		smartchain	LOOT	0x030d1aed728eeda26a07dbeb0e54a1e785822069f24d44ec6f00cf7de77c239e
		smartchain	MBOX	0xf1cf8e837fca54b3e86dd837085ecf85527261a9dc93acfc7b56aa833710f00b
		smartchain	MULTI	0xcba2c48aee26b22f3fb333082e884f31df6ba8b59a6d791b9a4a7c7b610536fa
		smartchain	POLC	0x30a1d84ccf4559d4fe535efaca32bb20975972fb856061b9367c52a7dd345308
		smartchain	PORTO	0xb88cba0fc113f119d6830c33d9f56a882d7972edf72d6d3f25a924d9d2d01d6a
		smartchain	PSG	0xcfb4d7b43fce3be58e148b64224f3af62ba03f98723b83a93021dda874a7aa3c
		smartchain	RACA	0x8d1be6af7b0f77cf950fdd8f1c6328763e6a8961fc26ef74cba7bbfebbca81f9
		smartchain	REEF	0x50f701ada157ab8f88535d978348ebb09ffb7fbe2aaa1a17a5eb21e58f6f18b1
		smartchain	REVO	0xef5ffe2092fc9a264e52efd6414c4b1ca80922b42785b097067b3fb8566d198e
		smartchain	SANTOS	0x8d046ebc7c0c120cf7d0ccd4db810503c7a915cb9f6b8b574a30f6f0831942e7
		smartchain	SFP	0x8651f485d74d12427838bda461611cd2156c0742d62da41e1b858633a9443c43
		smartchain	STG	0xcbfbcad264a91feb9241088deb18d1cfc3f24ebca9eee2f4360dc649e0326602
		smartchain	SU	0x681cdd3a9d522defdd22997952f05ef1072dcf58d64866be627812b1ed8ba06f
		smartchain	TKO	0xd6609dd01bdb73f687412592346e0455763e3acb56f9a31b46f7c07fce4d7b09
		smartchain	TORN	0x3a0098aa8128e873de9ecf2860bb3d4e67e4310512851e1c7f3cd1bd3e9720b4
		smartchain	TRVL	0xfe1d5e0f53fa19e8b824ac7040df1c2aac4c596a61a2f277b9b904b2559457e5
		smartchain	TWT	0x07d3626a55774c08719cf711fe2cad06e4373492498e51a73c3ca431bcc03010
		smartchain	UNFI	0xefd0ba6f430c7e289aba0fa18566aa44e02da636196770d08c583065e99a3926
		smartchain	USDC	0x82fb29d7dbbb69d69e50fa32c7db3d2e2398ac198ec3785b57dba12bd05081b1
		smartchain	USDT	0x341831807a8424ddf2a543f904e59bfcf71d225f503113c00cb89e0bbd6ea229
		smartchain	VRGW	0x079b67944e6c5a0ef5b9792ae93271bc77a905121e3f2b2c54786c2f96b06981
		smartchain	WWY	0x0028af1baf96af3b5b9c681c9877219a770f5042068e409d5fc313778fec5bfc
		smartchain	XVS	0x371ca3e421ba0dec4d30ad0f2e9c7bd6416213e9657417f2921f2755fa28a1e4		`
	*/
	type CoinTransaction struct {
		Chain    string
		Coin     string
		Hash     string
		GasLimit uint64
	}

	lines := strings.Split(txlists, "\n")
	coinTransactions := make([]*CoinTransaction, 0)
	for _, line := range lines {
		items := strings.Split(line, "\t")
		if items[0] == "" || items[1] == "" {
			continue
		}
		coinTransactions = append(coinTransactions, &CoinTransaction{
			Chain:    items[0],
			Coin:     items[1],
			Hash:     items[2],
			GasLimit: 0,
		})
	}
	//
	ethClient := DefaultEthereumClient()
	for _, coinTransaction := range coinTransactions {
		hash := common.HexToHash(coinTransaction.Hash)
		receipt, err := ethClient.GetTransactionReceipt(context.Background(), hash)
		if err != nil {
			panic(err)
		}
		coinTransaction.GasLimit = (receipt.GasUsed/100 + 1) * 100
	}
	lines = make([]string, 0)
	for _, coinTransaction := range coinTransactions {
		line := fmt.Sprintf("update t_tokens set gas_limit = %d where chain = '%s' and symbol = '%s';",
			coinTransaction.GasLimit,
			coinTransaction.Chain,
			coinTransaction.Coin)
		lines = append(lines, line)
	}
	output := strings.Join(lines, "\n")
	fmt.Printf("%s\n", output)
}
