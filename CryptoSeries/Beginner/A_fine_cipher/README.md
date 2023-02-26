# A Fine Cipher

**[Return to main page](../)**

- [Part 1: It is Not Fine](#part-1-it-is-not-fine)
- [Part 2: Extending the Alphabet](#part-2-extending-the-alphabet)
- [Part 3: A Likely Solution](#part-3-a-likely-solution)

---

#### Practice Ciphertext: [JSON](./ciphertext/ciphertext.json) | [TXT](./ciphertext/ciphertext.txt)

##### _Note:_ The JSON is in the format [{a<sup>-1</sup>, b, ciphertext}] and the text document is a tab deliminated list for each row formatted as "a<sup>-1</sup> b ciphertext"

---

## Part 1: It is Not Fine

![It's fine meme](https://i.kym-cdn.com/entries/icons/mobile/000/018/012/this_is_fine.jpg)

I have this encrypted text, and I've been told its a fine text. Now that I think about it, I think they meant affine text. Luckily, it seems to only be using the english alphabet, ignoring casing, and ignores punctuation and spaces. Also, lucky for you, I have been playing around with the text, and I would have to rate this text a 5 through a 15. Please help me make this affine text a fine text!

**The fine text:** _odi vade ro vchi rpis qhh, odi vade ro tadx rpis, odi vade ro bvade rpis qhh qdx ad rpi xqvwdigg badx rpis._

**Hints:**

<details>
<summary>A misleading rating</summary>
The rating range of 5 through 15 refers to the value of a used to encrypt the text. a<sup>-1</sup> is different but related to this value.
</details>

<details>
<summary>Not much to check</summary>
The number of combinations you need to check is less than 500! However, to get to this low number of checks, something has to be done to limit either a<sup>-1</sup> or b. Know that gcd(a, n) must equal 1, and that a(mod(n)) * a<sup>-1</sup>(mod(n)) = 1(mod(n))
</details>

## Part 2: Extending the Alphabet

![ASCII alphabet](https://res.cloudinary.com/practicaldev/image/fetch/s--2xoVYXR3--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/gcsd9q3utce801qbfghq.jpg)

What is this! Apparently, the alphabet can include more than the english alphabet!!! Ugh, I can't believe this. I need you help. I implore you help help me with decrypting this text. However, I'm pretty sure this text is rated a 1 to a 10. However, it now seems to include punction!

**The new text:** _!dn wnfoz o?arwxro ?k f jreend dnnop fko znwwj ? ln;eo kn! !wfare in!, fko ir nkr !wfarerwp enkx ? z!nno fko ennbro ondk nkr fz ufw fz ? ln;eo !n d,rwr ?! irk! ?k !,r ;korwxwnd!,: !,rk !nnb !,r n!,rwp fz -;z! fz uf?wp fko ,fa?kx qrw,fqz !,r ir!!rw lef?hp irlf;zr ?! dfz xwfzzj fko dfk!ro drfw: !,n;x, fz unw !,f! !,r qfzz?kx !,rwr ,fo dnwk !,rh wrfeej fin;! !,r zfhrp fko in!, !,f! hnwk?kx rt;feej efj ?k erfarz kn z!rq ,fo !wnoork ieflbs n,p ? brq! !,r u?wz! unw fkn!,rw ofjv jr! bknd?kx ,nd dfj erfoz nk !n dfjp ? on;i!ro ?u ? z,n;eo rarw lnhr iflbs ? z,fee ir !ree?kx !,?z d?!, f z?x, znhrd,rwr fxrz fko fxrz ,rklr. !dn wnfoz o?arwxro ?k f dnnop fko ?— ? !nnb !,r nkr erzz !wfarero ijp fko !,f! ,fz hfor fee !,r o?uurwrklrs_

**Hints:**

- The alphabet used was "abcdefghijklmnopqrstuvwxyz,.!?:;-\_"

<details>
<summary>A properly defined function</summary>
If you haven't already, your decode function is recommended to look like the following
<code>decode(a, b, c, n) -> int</code> where a is a<sup>-1</sup>, b is b, c is the position of the ciphertext character in the alphabet (zero-indexed), and n is the size of the alphabet. It might help to have the function return the position of the character instead of trying to decode the entire string within one function.
</details>

## Part 3: A Likely Solution

***Encouraged to work with a partner or group**

Now that you are able to solve affine ciphers with any alphabet. It would be really cool if there was a way to rank the brute force guesses based on which ones are readable in English, and preferably with the correct answer ranked number one!

The person/group who submits the most readable, fastest, and most accurate solution will have their solution added to ctf-tools.
