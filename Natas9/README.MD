Natas9


login:natas9
password:W0mMhUcRRnG8dcghE4qvk3JA9lGt8nDl


1)На главной странице есть поле ввода через которое мы можем осузествлять поиск.Есть возможность посмотреть исходный код (View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/173111079-fd4f8945-5553-4006-836c-5f1600c1164e.png)
2) Если посмотреть исходный код, то можно увидить такой код


    if($key != "") {
    passthru("grep -i $key dictionary.txt");
    }
    
Отсюда можно увидить что поиск происходит в файле dictionary.txt с помощью утилиты grep. Так же видно что никакой фильтрации нашего ввода нет. 
С помощью grep можно осуществлять поиск сразу по нескольким файлам, мы знаем что все пароли хранятся в /etc/natas_webpass/natas<number>, в нашем 
случае пароль должен находится в /etc/natas_webpass/natas10.
  
3)Что если отправить 1 /etc/natas_webpass/natas10 по итогу запрос превратится в 
  
    grep -i 1 /etc/natas_webpass/natas10 dictionary.txt
  
То есть grep будет искать цифру 1 в 2 файлах в dictionary.txt и etc/natas_webpass/natas10 и если данная цифра есть пароле то он нам вернется.
  ![изображение](https://user-images.githubusercontent.com/54982244/173113678-dde34d61-fc23-4945-9a4c-e06927a87fef.png)

2 вариант. 
Так как нет никакой фильтрации пользовательского ввода, то можно использовать ; для выполнения нескольких bash команд одновременно.
Что если отправить ;cat /etc/natas_webpass/natas10, по итогу запрос превратится в 
  
    grep -i ;cat /etc/natas_webpass/natas10 dictionary.txt
  
Команда grep не выполнится, но cat выведит содержимое обоих файлов.
 ![изображение](https://user-images.githubusercontent.com/54982244/173114367-37bc99fd-f65a-41da-84f3-16a555e9786f.png)
