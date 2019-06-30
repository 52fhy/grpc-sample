<?php

use Grpc\ChannelCredentials;
use Sample\Model\User;
use Sample\Model\UserList;
use Sample\Model\GreeterClient;

ini_set("display_errors", true);
error_reporting(E_ALL);
require_once "autoload.php";
$user = new User();
$user->setId(1)->setName("test");

$client = new GreeterClient("192.168.99.1:50051", [
    'credentials' => ChannelCredentials::createInsecure(), //不加密
//    'timeout' => 3000000,
]);

//分别是响应、状态对象
list($reply, $status) = $client->SayHello($user)->wait();

if (!$reply) {
    echo json_encode($status);
    return;
}

//序列化为string
echo $reply->serializeToJsonString(true) . PHP_EOL;
echo $reply->getErrCode() . PHP_EOL; //errCode
echo $reply->getErrMsg() . PHP_EOL; //errMsg

//data
foreach ($reply->getData() as $key => $value) {
    echo $key . "-" . $value . PHP_EOL;
}
