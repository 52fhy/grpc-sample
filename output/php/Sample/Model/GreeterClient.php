<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Sample\Model;

/**
 */
class GreeterClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Sends a greeting
     * @param \Sample\Model\User $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function SayHello(\Sample\Model\User $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/Sample.Model.Greeter/SayHello',
        $argument,
        ['\Sample\Model\Response', 'decode'],
        $metadata, $options);
    }

}
