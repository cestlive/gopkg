## Code snapshots
### shutdown

```go
// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := s.HttpServer.Shutdown(ctx); err != nil {
				l.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 grpc server
		func() {
			if s.GrpcServer != nil {
				s.GrpcServer.GracefulStop()
			}
		},

		// 关闭 db
		func() {
			if s.DB != nil {
				if err := s.DB.DbClose(); err != nil {
					l.Error("db close err", zap.Error(err))
				}
			}
		},

		// 关闭 Storage
		func() {
			if s.Storage != nil {
				if err := s.Storage.Close(); err != nil {
					l.Error("storage close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					l.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 Trace
		func() {
			if s.Trace != nil {
				// Do not make the application hang when it is shutdown.
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				if err := s.Trace.Shutdown(ctx); err != nil {
					l.Error("trace close err", zap.Error(err))
				}
			}
		},


```
