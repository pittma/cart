/*
Map your URIs.

Cart uses a [trie](http://en.wikipedia.org/wiki/Trie) based URI lookup algorithm.  This means that routing performance is independent from the number of routes.  When using a traditional URI dispatcher, worst case scenario time performance is O(n) where n is the total number of routes.  Using a trie, routes are dispatched at worst O(m) where m is the length of the longest URI.
*/
package cart
