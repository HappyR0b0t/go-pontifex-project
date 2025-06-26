# Description

Pontifex is an algorithm for ciphering and deciphering text messages, described in Neal Stephenson's novel "Cryptonomicon", created by Bruce Schneier. More widely known as Solitaire cipher, because it uses a deck of playing cards. Deck consits of 52 classic cards, plus two Jokers. This ciphering algorithm can be performed with physical deck of cards, but I decided to make a program for study puprposes and fun. 

For this algorithm, to be able to cipher a text message, it should consist of latin characters only (a-z), no symbols! The case of characters is irrelevant. Ciphered and deciphered message will be a string consisting of upper case characters (A-Z) grouped in 5 letters, for example:

    "Do not use pc" after ciphering will turn into a random set of characters that could look like this - "ZFQIK LCOUA"
    "ZFQIK LCOUA" after deciphering will turn into "DONOT USEPC". You get the idea. This grouping by 5 characters is just a common convention in cryptography.

It is important to mention, that in order to sucssesfully decipher a message, the same deck should be used. The deck is a text containing 54 "words" of "cards" in following format <suit>-<rank>. Suits are classical: clubs, diamond, hearts and spades. Ranks are as follows: Ace, 2, 3 ... King and two Jokers: JA and JB. If you want another person to be able to decipher your message, you should provide a ciphered text and the deck you used to cipher it (in its starting state).

Have fun. 

## How to install and run

Go version - go1.23.0 windows/amd64

Clone a repo to your local machine:

`git clone git@github.com:HappyR0b0t/go-pontifex-project.git`

To run the algorithm:

`go run .`

(from the root folder of local repo)

To run tests:

`go test ./...`

(from the root folder of local repo)


## How to use

There are two important entities that are used as an input for algorithm to work:
    - a plain text
    - a deck in the form of text

