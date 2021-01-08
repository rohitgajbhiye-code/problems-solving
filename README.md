# problems-solving

1. Pascal Triangle 
2. Perfect reversible string - A simple solution for this problem is to generate all possible substrings of ‘st’ and check if their reverse exist in the ‘str’ linearly.
An efficient solution for this problem is based on the fact that reverse of all substrings of ‘str’ will exist in ‘str’ if and only if the entire string ‘str’ is palindrome. We can justify this fact by considering the whole string, a reverse of it will exist only if it is palindrome. And if a string is palindrome, then all reverse of all substrings exist.
