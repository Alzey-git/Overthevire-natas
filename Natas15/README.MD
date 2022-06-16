Natas15

login: natas15
password: AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J

1) На главной странице есть поле ввода. Можно проверить существует ли пользователь в системе введя имя пользователя. Есть возможность посмотреть исходный код (View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/174034648-4ff8c566-c1a6-407c-979f-f4b687eca84e.png)

Если пользователь существует, вывод - This user exists. (Например username=natas16)

Если пользователь не существует то вывод - This user doesn't exist.
![изображение](https://user-images.githubusercontent.com/54982244/174035378-ae241435-705b-4c86-9984-288d951764dc.png)

2) Исходный код. В коментарие видно что таблица users к которой и формируется запрос состоит из 2 полей, username и password. Фильтрации нашего ввода нет, возможна sql инъекция, только вот и вывода нет, значит это будет слепая sql инъекия.


        <?
        /*
        CREATE TABLE `users` (
          `username` varchar(64) DEFAULT NULL,
          `password` varchar(64) DEFAULT NULL
        );
        */

        if(array_key_exists("username", $_REQUEST)) {
            $link = mysql_connect('localhost', 'natas15', '<censored>');
            mysql_select_db('natas15', $link);

            $query = "SELECT * from users where username=\"".$_REQUEST["username"]."\"";
            if(array_key_exists("debug", $_GET)) {
                echo "Executing query: $query<br>";
            }

            $res = mysql_query($query, $link);
            if($res) {
            if(mysql_num_rows($res) > 0) {
                echo "This user exists.<br>";
            } else {
                echo "This user doesn't exist.<br>";
            }
            } else {
                echo "Error in query.<br>";
            }

            mysql_close($link);
        } else {
        ?> 

3) В поле ввода введем такую строку - natas16" AND password LIKE BINARY 'W%';-- - по итогу наш запрос превратистя в 
    
        SELECT * from users where username="natas16" AND password LIKE BINARY 'W%';-- -
        
То есть пройдет запрос к таблице users, где username = natas16. password LIKE BINARY 'W%' вернет true только в том случае если первая буква
пароля будет равна букве W (С учетом регистра.) Если все верно и запрос в целом возвращает true, на экран выведится This user exists. 
Если нет, то есть если первая буква пароля не W, на экран выведится This user doesn't exist. И таким перебором необходимо подобрать пароль.

![изображение](https://user-images.githubusercontent.com/54982244/174038082-96f29690-132a-431c-9d0b-1398801bb5f0.png)
![изображение](https://user-images.githubusercontent.com/54982244/174038054-bbb79acd-bb43-4426-88ea-8be3dcebbdb3.png)

4) Скрипт для перебора. Таким образом скрипт посимвольно перебирает варианты пока не получит пароль.
![изображение](https://user-images.githubusercontent.com/54982244/174039199-5a16fb35-65b7-454b-92d7-a0ea9afdf216.png)
![изображение](https://user-images.githubusercontent.com/54982244/174040390-95905fa3-92e1-444c-ac1f-a4b4da6c3fc2.png)


