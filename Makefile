default:
		go run -ldflags "-X main.CommitHash=`git rev-parse HEAD`" -race ./*.go
rice:
		rice embed-go -i ./api && go run -ldflags "-X main.CommitHash=`git rev-parse HEAD`" -race ./*.go