Natas10

login:natas10
password:nOpp1igQAkUzaI1GUUjzn1bFVj7xCNzu

1)На главной странице есть поле ввода через которое мы можем осуществлять поиск. Так же написано что есть фильтрация некоторых символов.
Есть возможность посмотреть исходный код (View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/173180066-8aea2f86-699f-4a28-983b-9a2263065240.png)


2)В исходном коде видим такой код


    <?
    $key = "";

    if(array_key_exists("needle", $_REQUEST)) {
        $key = $_REQUEST["needle"];
    }

    if($key != "") {
        if(preg_match('/[;|&]/',$key)) {
            print "Input contains an illegal character!";
        } else {
            passthru("grep -i $key dictionary.txt");
        }
    }
    ?>
Видно что функция preg_match фильтрует такие символы как ;|& ,поэтому не получится использовать ; для объединения нескольких bash команд. 

3)С помощью grep можно осуществлять поиск сразу по нескольким файлам, мы знаем что все пароли хранятся в /etc/natas_webpass/natas, 
в нашем случае пароль должен находится в /etc/natas_webpass/natas11. Если мы отправим 1 /etc/natas_webpass/natas11 в поле ввода, наш итоговый запрос превратится в

    grep -i 1 /etc/natas_webpass/natas11 dictionary.txt 
    
Утилита grep будет искать цифру 1 в 2 файлах, /etc/natas_webpass/natas11 и dictionary.txt и если цифра 1 присутствует в нашем пароле то он вернется, если нет
то будет необходимо перебирать символы.
Нам повезло и цифра 1 присутствует в парооле.
![изображение](https://user-images.githubusercontent.com/54982244/173180209-4962064b-bde5-4701-a274-7049dad25d84.png)

