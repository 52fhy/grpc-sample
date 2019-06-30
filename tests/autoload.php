<?php 

//自动加载
spl_autoload_register(function($class_name)
{
    if (class_exists($class_name)) {
        return;
    }

    $split = explode('\\', $class_name);
    if($split[0] != 'Sample'){
        return;
    }

    $classFile = dirname(__DIR__). '/output/php/' . join($split, "/"). ".php";
    // echo $classFile.PHP_EOL;
    if (file_exists($classFile)) {
        require_once $classFile;
    }

    return;
});

include_once "vendor/autoload.php";