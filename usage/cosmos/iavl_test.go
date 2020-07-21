package cosmos

import (
	"fmt"
	"github.com/tendermint/iavl"
	"github.com/tendermint/tm-db"
	"testing"
)

func TestIAVLTree(t *testing.T) {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		panic(err)
	}
	tree.Set([]byte("e"), []byte{5})
	tree.Set([]byte("d"), []byte{4})
	tree.Set([]byte("c"), []byte{3})
	tree.Set([]byte("b"), []byte{2})
	tree.Set([]byte("a"), []byte{1})

	fmt.Printf("tree root hash %x\n", tree.Hash())

	rootHash, version, err := tree.SaveVersion()
	if err != nil {
		panic(err)
	}

	fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
	fmt.Println(tree.String())
	fmt.Printf("tree root hash %x\n", tree.Hash())
}

func TestIAVLTreePersistence(t *testing.T) {
	{
		mydb, err := db.NewGoLevelDB("iavl-test", "./")
		if err != nil {
			panic(err)
		}
		tree, err := iavl.NewMutableTree(mydb, 0)
		if err != nil {
			panic(err)
		}
		tree.Set([]byte("e"), []byte{5})
		tree.Set([]byte("d"), []byte{4})
		tree.Set([]byte("c"), []byte{3})
		tree.Set([]byte("b"), []byte{2})
		tree.Set([]byte("a"), []byte{1})

		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())

		mydb.Close()
	}
	{
		mydb, err := db.NewGoLevelDB("iavl-test", "./")
		if err != nil {
			panic(err)
		}
		tree, err := iavl.NewMutableTree(mydb, 0)
		if err != nil {
			panic(err)
		}
		tree.LoadVersion(1)
		tree.Set([]byte("f"), []byte{6})

		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())

		mydb.Close()
	}
}

func TestIAVLTreePersistence_NotSave(t *testing.T) {
	{
		mydb, err := db.NewGoLevelDB("iavl-test-not-save", "./")
		if err != nil {
			panic(err)
		}
		tree, err := iavl.NewMutableTree(mydb, 0)
		if err != nil {
			panic(err)
		}
		tree.Set([]byte("e"), []byte{5})
		tree.Set([]byte("d"), []byte{4})
		tree.Set([]byte("c"), []byte{3})
		tree.Set([]byte("b"), []byte{2})
		tree.Set([]byte("a"), []byte{1})

		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())

		mydb.Close()
	}
	{
		mydb, err := db.NewGoLevelDB("iavl-test-not-save", "./")
		if err != nil {
			panic(err)
		}
		tree, err := iavl.NewMutableTree(mydb, 0)
		if err != nil {
			panic(err)
		}
		tree.LoadVersion(1)
		tree.Set([]byte("f"), []byte{5})

		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())

		mydb.Close()
	}
}

func TestIAVLTreePersistence_MemDB(t *testing.T) {
	{
		tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
		if err != nil {
			panic(err)
		}
		tree.Set([]byte("e"), []byte{5})
		tree.Set([]byte("d"), []byte{4})
		tree.Set([]byte("c"), []byte{3})
		tree.Set([]byte("b"), []byte{2})
		tree.Set([]byte("a"), []byte{1})

		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())
	}
	{
		tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
		if err != nil {
			panic(err)
		}

		tree.LoadVersion(1)
		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		fmt.Println(tree.String())
		fmt.Printf("tree root hash %x\n", tree.Hash())
	}
}

func TestProof(t *testing.T) {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		panic(err)
	}
	tree.Set([]byte("e"), []byte{5})
	tree.Set([]byte("d"), []byte{4})
	tree.Set([]byte("c"), []byte{3})
	tree.Set([]byte("b"), []byte{2})
	tree.Set([]byte("a"), []byte{1})

	tree.SaveVersion()
	rootHash, version, err := tree.SaveVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("saved version %v with root hash %x\n", version, rootHash)

	key := []byte("d")
	value, proof, err := tree.GetWithProof(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("value of %s is %x\n", key, value)

	err = proof.Verify(tree.Hash())
	if err != nil {
		panic(err)
	}
}


func TestProof1(t *testing.T) {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		panic(err)
	}
	tree.Set([]byte("e"), []byte{5})
	tree.Set([]byte("d"), []byte{4})
	tree.Set([]byte("c"), []byte{3})
	tree.Set([]byte("b"), []byte{2})
	tree.Set([]byte("a"), []byte{1})

	tree.SaveVersion()
	rootHash, version, err := tree.SaveVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("saved version %v with root hash %x\n", version, rootHash)

	key := []byte("f")
	value, proof, err := tree.GetWithProof(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("value of %s is %x\n", key, value)

	err = proof.Verify(tree.Hash())
	if err != nil {
		panic(err)
	}
}

func TestProof2(t *testing.T) {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		panic(err)
	}
	tree.Set([]byte("e"), []byte{5})
	tree.Set([]byte("d"), []byte{4})
	tree.Set([]byte("c"), []byte{3})
	tree.Set([]byte("b"), []byte{2})
	tree.Set([]byte("a"), []byte{1})

	for i := 0;i < 10;i ++ {
		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}
		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
	}
}


func TestProof3(t *testing.T) {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		panic(err)
	}
	for i := 0;i < 10;i ++ {
		tree.Set([]byte("e"), []byte{5})
		tree.Set([]byte("d"), []byte{4})
		tree.Set([]byte("c"), []byte{3})
		tree.Set([]byte("b"), []byte{2})
		tree.Set([]byte("a"), []byte{1})

		rootHash, version, err := tree.SaveVersion()
		if err != nil {
			panic(err)
		}
		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
	}
}
