> Sample projects used to learn GO

NOTE: 
- A sentinel error is pretty limited in what it can convey. 
  It can’t contain any dynamic information that we’ve learned at run-time.

### Good comment practice :
- Say what your package and function does and why if needed 
__NOTE__ : when you pubish your project to a hosting site such as GitHub,  
your comments will be automatically turned into browsable documentation on the pkg.go.dev site. 

### “red, green, refactor” cycle.
1. Write test for function that does not even exits . 
2. Run it and get compiler error for undefined function
3. Write the function just enough to make test compile and fail
4. check that the failure message is accurate and informative.
    - in test case where there should be no error check for err == nil
    -  and when writing invalid input , check for err != nil
5. Write the minimum code necessary to make the test pass.
6. Optionally, tweak and improve the code, preferably without breaking the test.
7. COMMIT !


### NOTES:
- For declaration of variables, if you want the default value, use var else use := 
- Slices in Go are ordered but not maps




