package main

func main() {
	user := person{
		firstName: "Iago",
		lastName:  "Maia",
		contactInfo: ContactInfo{
			zipCode: 123321,
			email:   "iago@email.com",
		},
	}
	user.printName()
	user.updateName("Jimmy")
	user.printName()
}
