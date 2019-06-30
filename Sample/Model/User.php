<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: User.proto

namespace Sample\Model;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>Sample.Model.User</code>
 */
class User extends \Google\Protobuf\Internal\Message
{
    /**
     *主键id
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    private $id = 0;
    /**
     *用户名
     *
     * Generated from protobuf field <code>string name = 2;</code>
     */
    private $name = '';
    /**
     *头像
     *
     * Generated from protobuf field <code>string avatar = 3;</code>
     */
    private $avatar = '';
    /**
     *地址
     *
     * Generated from protobuf field <code>string address = 4;</code>
     */
    private $address = '';
    /**
     *手机号
     *
     * Generated from protobuf field <code>string mobile = 5;</code>
     */
    private $mobile = '';
    /**
     *扩展信息
     *
     * Generated from protobuf field <code>map<string, string> ext = 6;</code>
     */
    private $ext;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *          主键id
     *     @type string $name
     *          用户名
     *     @type string $avatar
     *          头像
     *     @type string $address
     *          地址
     *     @type string $mobile
     *          手机号
     *     @type array|\Google\Protobuf\Internal\MapField $ext
     *          扩展信息
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User::initOnce();
        parent::__construct($data);
    }

    /**
     *主键id
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     *主键id
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

    /**
     *用户名
     *
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     *用户名
     *
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     *头像
     *
     * Generated from protobuf field <code>string avatar = 3;</code>
     * @return string
     */
    public function getAvatar()
    {
        return $this->avatar;
    }

    /**
     *头像
     *
     * Generated from protobuf field <code>string avatar = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setAvatar($var)
    {
        GPBUtil::checkString($var, True);
        $this->avatar = $var;

        return $this;
    }

    /**
     *地址
     *
     * Generated from protobuf field <code>string address = 4;</code>
     * @return string
     */
    public function getAddress()
    {
        return $this->address;
    }

    /**
     *地址
     *
     * Generated from protobuf field <code>string address = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setAddress($var)
    {
        GPBUtil::checkString($var, True);
        $this->address = $var;

        return $this;
    }

    /**
     *手机号
     *
     * Generated from protobuf field <code>string mobile = 5;</code>
     * @return string
     */
    public function getMobile()
    {
        return $this->mobile;
    }

    /**
     *手机号
     *
     * Generated from protobuf field <code>string mobile = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setMobile($var)
    {
        GPBUtil::checkString($var, True);
        $this->mobile = $var;

        return $this;
    }

    /**
     *扩展信息
     *
     * Generated from protobuf field <code>map<string, string> ext = 6;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getExt()
    {
        return $this->ext;
    }

    /**
     *扩展信息
     *
     * Generated from protobuf field <code>map<string, string> ext = 6;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setExt($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->ext = $arr;

        return $this;
    }

}

