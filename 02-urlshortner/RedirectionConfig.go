package main 

type RedirectRule struct {
  Slug string
	Url  string
}

type RedirectionConfig struct {
	Redirects []RedirectRule
}

