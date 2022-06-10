Natas6

login:natas6
password:aGoY4q2Dc6MgDq4oL4YtoKtyAg9PeHa1

1)На главной странице видим строку ввода в которую просят отправить секретное сообщение что бы получить пароль для следующего уровня.
![изображение](https://user-images.githubusercontent.com/54982244/173066679-6c7bd0c5-18f0-4a8c-a491-c01dc2a4f200.png)
2)У нас есть возможность посмотреть исходный php код (View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/173066745-5b13621f-d7cd-496d-a327-909143f450f5.png)
3)Судя поисходному php коду, идет простое сравнение введеной нашей строки с $secret, сам же $secret подгружается из файла includes/secret.inc, проверим, 
сможем ли мы напрямую прочесть файл includes/secret.inc. http://natas6.natas.labs.overthewire.org/includes/secret.inc
![изображение](https://user-images.githubusercontent.com/54982244/173067295-bbeb1505-6cf1-4649-b732-f52a4f8baaaa.png)
4)При переходе открывается пустая страница, переходим в исходный код и видим 

![изображение](https://user-images.githubusercontent.com/54982244/173067496-876a9019-0c45-4d35-b95d-5ee615f7320e.png)

Забираем наше секретное слово.

5)Отправляем секретное слово, и забираем пароль для следующего уровня.
![изображение](https://user-images.githubusercontent.com/54982244/173067624-455622dc-8968-4adf-b8b3-3dcfb541d85f.png)
