package config

func (init *TInit) InitProxies() {
	init.app.ForwardedByClientIP = true
	init.app.SetTrustedProxies([]string{
		"localhost",
		"127.0.0.1",
	})
}
