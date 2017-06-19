package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stevedonovan/luar"
	"fmt"
	"github.com/stephenlyu/goformula/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
)

var _ = Describe("LUAR", func() {
	It("test", func () {
		const test = `
for i = 1, 3 do
		print(msg, i)
end
print(user)
print(user.Name, user.Age)
`

		type person struct {
			Name string
			Age  int
		}

		L := luar.Init()
		defer L.Close()

		user := &person{"Dolly", 46}

		luar.Register(L, "", luar.Map{
			// Go functions may be registered directly.
			"print": fmt.Println,
			// Constants can be registered.
			"msg": "foo",
			// And other values as well.
			"user": user,
		})

		L.DoString(test)
	})
})

var _ = Describe("VectorLua", func() {
	It("test", func () {
		const test = `

`
		L := luar.Init()
		defer L.Close()

		//v := function.Vector([]float64{1, 2, 3, 4, 5})

		luar.Register(L, "", luar.Map{
			// Go functions may be registered directly.
			"print": fmt.Println,
			// Constants can be registered.
			//"v": v,
			"Vector": function.Vector,
			"Close": stockfunc.CLOSE,
			"RVector": stockfunc.RecordVector,
		})

		L.DoString(test)
	})
})

var _ = Describe("Account", func() {
	It("test", func () {
		L := luar.Init()
		defer L.Close()

		//v := function.Vector([]float64{1, 2, 3, 4, 5})

		luar.Register(L, "", luar.Map{
			// Go functions may be registered directly.
			"print": fmt.Println,
			// Constants can be registered.
			//"v": v,
			"Vector": function.Vector,
			"Close": stockfunc.CLOSE,
			"RVector": stockfunc.RecordVector,
		})

		L.DoFile("account.lua")

		//b = Account.new(Account, {value=1000})
		//--b.display(b)


		L.GetGlobal("Account")
		println(L.GetTop())
		L.GetField(-1, "new")
		println(L.GetTop())

		L.PushValue(-2)
		println(L.GetTop())
		L.NewTable()
		println(L.GetTop())
		L.Call(2, 1)
		println(L.GetTop())

		L.Remove(1)
		println(L.GetTop())

		L.GetField(-1, "display")
		L.PushValue(-2)
		L.Call(1, 0)
		println(L.GetTop())

	})
})
