# NetSpk

## Start service exec
Clone repo and run the following to start listening on port 9000
```
chmod +x ./followerservice
./followerservice
```

#### Alternatively 
Run the package from the repo root
```
go run ./cmd/followerservice  
```

## Create a user
```
grpcurl -plaintext -d '{"email":"you@email.com", "screenName":"player-x"}' localhost:9000 gen.FollowerService.CreateUser
```

## Follow a user
Using UUIDs returned from the above...
```
grpcurl -plaintext -d '{"id":"_GUID_", "followId":"_GUID_"}' localhost:9000 gen.FollowerService.FollowUser
```
