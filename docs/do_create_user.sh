# for convienince 
grpcurl -plaintext -d '{"email":"toms email", "screenName":"toms screen name"}' localhost:9000 gen.FollowerService.CreateUser