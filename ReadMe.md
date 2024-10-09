

# Dummy Web

This is a dummy domain module

It is a resource for 

workspace example 
or 
replace

For workspace see (workspace repo)[https://github.com/mkdirJava/dummy_workspace]

---

As for replace

This service uses (dummy_domain)[https://github.com/mkdirJava/dummy_domain].
The dummy domain exposes a service that returns a string. 
Suppose we need to develop this and want to change the return text.
We could do it in a number of ways but we want to not impact other developers.
So we

1. Git clone (dummy_domain) [https://github.com/mkdirJava/dummy_domain]
2. Alter this go.mod to replace a local version
  
    at the root of this project Run

    go mod edit -replace github.com/mkdirjava/dummy_domain=../dummy_domain
    do mod tidy

    The above should point the module to use local dummy_domain

3. Alter the dummy_domain just cloned and alter local test in server_test.go
4. Ensure the test pass
5. Push both the Dummy_domain as well as the dummy_web for the release process
6. Note the go mod needs to be replaced in the release process


To remove the replace, the easiest way is to remove the require and replace sections in go.mod and run 

  go mod tidy

go mod tidy will read the modules imports and import then in the go.mod as well as get the dependencies checksums and produce a go.sum. Always check both of these in
