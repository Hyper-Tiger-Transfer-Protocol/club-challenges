# Cryptography Challenges
Every challenge within this section is designed to build and test various aspects of cryptographic knowledge. As the challenges grow in number, so too shall this document. This document will contain the list of challenges and resources to learn about the principle(s) being tested. Each challenge will have its own **README** as well. Therein you will find the outline of the challenge, the parts it has, and hints for the challenge. Later on, these **READMEs** will also contain the number of points the challenges are worth.

**NOTE:** Each part of a challenge is worth points.

---
# Table of Contents
- [Challenges](#challenges)
  - [Beginner Challenges](#beginner-challenges)
      - [A Fine Cipher](#a-fine-cipher)
      - [CAESARRRRRRRRR!](#caesarrrrrrrrr)
      - [On The Fence](#on-the-fence)
      - [Rank And File](#rank-and-file)
  - [Intermediate Challenges](#intermediate-challenges)
  - [Difficult Challenges](#difficult-challenges)
  - [Master Challenges](#master-challenges)
---

# Challenges
Below are the challenges and some recommended resources to help in solving the, barring hints. The linked plaintext documents are the strings used in the provided ciphertext documents. The provided ciphertexts are not a part of the final flag and are for testing purposes.

#### Plaintext: [JSON](./plaintext/plaintext.json) | [TXT](./plaintext/plaintext.txt)

In addition, examples for how to use json in some programming languages are given [here](./JSON_import_examples/).

---

## Beginner Challenges

#### A Fine Cipher
##### [What is an affine cipher?](https://en.wikipedia.org/wiki/Affine_cipher)
##### [Challenge README](./A_fine_cipher/README.md)
This challenge covers affine ciphers, a type of substitution cipher. In general, the encrypted letters follow the following pattern:
> Let p be the plaintext character and c be the ciphertext character of the alphabet A.
> c = (a*p + b)mod(n), where a and b are scalars with a != 0, n is the size of the alphabet, |A|, and p and c are abstracted to their position in within A indexed at 0.
>
> For example, let A be the english alphabet, |A| = n = 26, a = 3, b = 1, p = 'c' = 2.
> c = (a*p + b)mod(n) = (3*2 + 1)mod(26) = (6 + 1)mod(26) = (7)mod(26) = 5 = 'h'

If you want to learn more about why affine works as a cipher and why the decryption function for it requires a and n to be coprime, I recommend you learn about the modulus function and congruence classes.

#### CAESARRRRRRRRR!
##### [What is a rotation cipher?](https://en.wikipedia.org/wiki/Caesar_cipher)
##### [Challenge README](./CAESARRRRRRRRRR!/README.md)

#### On The Fence
##### [What is a transposition cipher?](https://en.wikipedia.org/wiki/Transposition_cipher)
##### [Challenge README](./On_the_fence/README.md)

#### Rank And File
##### [What is a transposition cipher?](https://en.wikipedia.org/wiki/Transposition_cipher)
##### [Challenge README](./Rank_and_File/README.md)

---
## Intermediate Challenges

---
## Difficult Challenges

---
## Master Challenges
---