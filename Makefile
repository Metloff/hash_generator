default:
		go run -ldflags "-X main.CommitHash=`git rev-parse HEAD`" -race ./*.go