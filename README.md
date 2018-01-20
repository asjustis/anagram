# anagram challenge

## Challenge

We have a message for you. But we hid it. 
Unless you know the secret phrase, it will remain hidden.

Can you write the algorithm to find it?

Here is a couple of important hints to help you out:
- An anagram of the phrase is: "poultry outwits ants"
- There are three levels of difficulty to try your skills with
- The MD5 hash of the easiest secret phrase is "e4820b45d2277f3844eac66c903e84be"
- The MD5 hash of the more difficult secret phrase is "23170acc097c24edb98fc5488ab033fe"
- The MD5 hash of the hard secret phrase is "665e5bcb0c20062fe8abaaf4628bb154"
Here is a list of english words, it should help you out.

## Solution

Most trivial way: try all permutations of all words in dictionary. But it will not work for us - too slow. 

Knowing that anagram is same letters in different order - let's get rid of words from the list that have any character NOT from the anagram phrase. Here we are, from around 100.000 words to just around 2700. 

Trying a straighforward approach to try all 3-word permutations now sounds feasible. So a simple nested loop of O(n^3) is completed in a few hours. The result is that easiest and medium difficulty md5 hashes are found. Perfect.

## Improvements

Thinking of the hardest hash, few ideas come to mind: maybe an apostrophe and some foreign letters (é instead of e, etc.) are used in the solution? Maybe it is just different amount of words used, rather than 3? Most likely.

Tried a simple recursive approach to this, yet it is too slow. Some improvement is a must to find a solution. Can it be just parallelized and solved on many computers? Can BFS instead of DFS approach would be enough? Can we find some more rules on how to discard many more cases without even trying them out? Well, it's a thought for the next time. 