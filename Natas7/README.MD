Natas7

login:natas7
password:7z3hEENjQtflzgnT29q7wAvMNfZdh0i9

1)На главной странице есть 2 вкладки, home и about
![изображение](https://user-images.githubusercontent.com/54982244/173069039-5631c82d-bd8d-46ea-9cb3-c226111e6185.png)
2)В исходном коде можно увидеть такую запись: "hint: password for webuser natas8 is in /etc/natas_webpass/natas8". Подсказка сообщающая путь где хранится 
пароль для следующего уровня.
![изображение](https://user-images.githubusercontent.com/54982244/173069264-7de0daf5-8d52-4f66-b7e8-121e3f6b6cbd.png)
3)При переходе по вкладкам home и about можжно заметить что переход осуществляется с помощью get параметра /index.php?page=home, в этом случае можно
проверить данный параметр на наличие LFI уязвимости, добавив вместо home такую полезную нагрузку page=../../../../../../../etc/passwd
(../ - переход на директорию ниже. С помощью множества ../ мы спускаемся в корневую директорию)

LFI - это возможность использования локальных системных файлов сервера. Уязвимость позволяет удаленному пользователю получить доступ с помощью специально
сформированного запроса к произвольным файлам на сервере, в том числе содержащую конфиденциальную информацию
Подробнее: https://insafety.org/lfi.php

И да, у нас получилось подгрузить вместо файла home, файл /etc/passwd
![изображение](https://user-images.githubusercontent.com/54982244/173069779-e8ee3ce5-2915-4eb2-8000-db783291d4ff.png)

4) Теперь вместо /etc/passwd попробуем получить доступ к файлу /etc/natas_webpass/natas8 добавив такую полузную нагрузку 
page=../../../../../../etc/natas_webpass/natas8
![изображение](https://user-images.githubusercontent.com/54982244/173070470-41cce5e1-f956-4003-86d0-722260a7949a.png)
5)Забираем пароль для следующего уровня.
