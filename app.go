package main

var config = Config{}

func init() {
	config.Read()
}

func main() {
	subject := "Get latest Tech News directly to your inbox"
	destination := "mohamed.labouardy@gmail.com"
	r := NewRequest([]string{destination}, subject)
	r.Send("templates/template.html", map[string]string{"username": "Conor"})
}
