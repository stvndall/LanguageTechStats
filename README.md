# Language Tech Stats

This is a repo that should help guide people to which language[s] to use for a new microservice project. 

I Do not know if I am always following best practices for a given language
Pull requests for when you :rage: are always welcome :tada:

## The Goal

This repository contains multiple languages, doing the exact same algorithm. But with different technologies.
This algorithm will do testing of

1. Speed of a given language
1. Readability of the code
1. How concurrency looks, and works if available
1. To as much as posible only use base language constructs
1. Guage how useful tech might be in your case
    * Current Tech that will be in repo
        * [gRPC](https://www.grpc.io) : a blazing fast rpc framework with cross language support
        * [Linkerd](https://www.linkerd.io) : A service mesh that will assist with communication of servers. And remove all the hard networking bits.

### algorithm

This algorithm is designed to test

1. Streams and basic functions
1. CPU time of simple tasks
1. CPU time of more complex tasks

|variable Name | Required variables |  Description | 
|----|----|----|
| _input_ |  | A given number from the caller >= 2 |
| _fibIdx_ | **_input_** | Given a number Find the index in fibonacci, or the previous index if not found |
| _words_  | **_fibIdx_** | Generate seeded list of words to the count of above index | 
| _max_ | **_words_** | Count the letter that appears the most | 
| _min_ | **_words_** | Count the letter that appears the least | 
| _factors_| **_fibIdx_** | find the [prime factors](https://www.mathsisfun.com/prime-factorization.html) | 
| _average_ | **_max_**,  **_min_**, **_factors_** | Find the average |
| | | return **_words_**, **_factors_**, **_average_** to caller |