//Welcome To My Ddos Attack Server (Site With IP)

func main() {
	workers := 100
	d, err := ddos.New("http://127.0.0.1:80", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://127.0.0.1:80")
	// Output: DDoS attack server: http://127.0.0.1:80
}
