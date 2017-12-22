beanstalk = Beanstalk::Pool.new([‘10.0.1.5:11300’])
loop do
  job = beanstalk.reserve
  puts job.body # prints “hello”
  job.delete
end



/ Produce jobs:

c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
id, err := c.Put([]byte("hello"), 1, 0, 120*time.Second)

/ Consume jobs:

c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
id, body, err := c.Reserve(5 * time.Second)
