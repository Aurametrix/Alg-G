/ Let's first create a client by providing adresses of nodes in the sandglass cluster
client, err := sg.NewClient(
    sg.WithAddresses(":7170"),
)
if err != nil {
    panic(err)
}
defer client.Close()

// Now we produce a new message
// Notice the empty string "" in the 3th argument, meaning let sandglass choose a random partition
err := client.Produce(context.Background(), "payments", "", &sgproto.Message{
    Value: []byte("Hello, Sandglass!"),
})
if err != nil {
    panic(err)
}
