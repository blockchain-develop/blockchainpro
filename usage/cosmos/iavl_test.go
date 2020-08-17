package cosmos

import (
	"fmt"
	"github.com/tendermint/iavl"
	"github.com/tendermint/tm-db"
	"sync"
	"testing"
	"time"
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

type StateTreeStableTest struct {
	tree *iavl.MutableTree
	exit chan bool
	lock sync.Mutex
}

func (this *StateTreeStableTest) updateTree() {
	for i := 0;i < 10000000;i ++ {
		//this.lock.Lock()
		this.tree.Set([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", i)))
		rootHash, version, err := this.tree.SaveVersion()
		//this.lock.Unlock()
		if err != nil {
			panic(err)
		}
		fmt.Printf("saved version %v with root hash %x\n", version, rootHash)
		time.Sleep(time.Millisecond * 1)
	}
	this.exit <- true
}

func (this *StateTreeStableTest) getProof() {
	for i := 0;i < 1000000;i ++ {
		//this.lock.Lock()
		key := []byte(fmt.Sprintf("%d", i))
		latestTree, err := this.tree.GetImmutable(this.tree.Version())
		if err != nil {
			continue
		}
		value, proof, err := latestTree.GetWithProof(key)
		hash := latestTree.Hash()
		//this.lock.Unlock()
		if err != nil {
			panic(err)
		}
		fmt.Printf("value of %s is %x\n", key, value)
		if proof != nil {
			err = proof.Verify(hash)
			if err != nil {
				panic(err)
			}
			fmt.Printf("proof verify successful!")
		}
		time.Sleep(time.Millisecond * 10)
	}
	this.exit <- true
}

func TestProof_MultiThread(t *testing.T) {
	mydb, err := db.NewGoLevelDB("iavl-multi-test", "./")
	if err != nil {
		panic(err)
	}
	tree, err := iavl.NewMutableTree(mydb, 0)
	if err != nil {
		panic(err)
	}

	treeTest := &StateTreeStableTest{
		tree : tree,
		exit: make(chan bool),
	}
	go treeTest.updateTree()
	go treeTest.getProof()

	<- treeTest.exit
	<- treeTest.exit
}
