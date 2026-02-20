# 1657. Determine if Two Strings Are Close

Две строки считаются близкими, если вы можете получить одну из другой, используя следующие операции:

- Операция 1: Поменять местами любые два существующих символа.
  
  Например, abcde -> aecdb
  
- Операция 2: Преобразовать все вхождения одного существующего символа в другой существующий символ, и сделать то же самое с другим символом.
  
  Например, aacabb -> bbcbaa (все a's превращаются в b's, и все b's превращаются в a's)
  
Вы можете использовать операции на любой строке столько раз, сколько захотите.

Даны две строки, word1 и word2, верните true, если word1 и word2 близки, и false в противном случае.


## Пример 1:

Input: word1 = "abc", word2 = "bca"

Output: true

Explanation: Вы можете получить word2 из word1 за 2 операции.
- Apply Operation 1: "abc" -> "acb"
- Apply Operation 1: "acb" -> "bca"


## Пример 2:

Input: word1 = "a", word2 = "aa"

Output: false

Explanation: Невозможно получить word2 из word1, или наоборот, за любое количество операций.


## Пример 3:

Input: word1 = "cabbba", word2 = "abbccc"

Output: true

Explanation: Вы можете получить word2 из word1 за 3 операции.
- Apply Operation 1: "cabbba" -> "caabbb"
- Apply Operation 2: "caabbb" -> "baaccc"
- Apply Operation 2: "baaccc" -> "abbccc"


## Ограничения:

- 1 <= word1.length, 
- word2.length <= 105
- word1 и word2 содержат только строчные английские буквы.