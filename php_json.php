<?php

function loadJson(){
	return file_get_contents('public_timeline.json');
}

function decode($json, $loop){
    echo '=== json decode ===' . "\n";
	$total_time = 0;
    foreach (range(0, $loop-1) as $i) {
        $start_point_m = microtime();
        $start = time();
        json_decode($json);
        $end_point_m = microtime();
        $end = time();
        $c = ($end - $start) + ($end_point_m - $start_point_m);
        $total_time += $c;
    }
    printf("%s call(s)/s \n", (1/($total_time/$loop)));
}

function encode($obj, $loop){
    echo '=== json encode ===' . "\n";
    $total_time = 0;
    foreach (range(0, $loop-1) as $i) {
        $start_point_m = microtime();
        $start = time();
        json_encode($obj);
        $end_point_m = microtime();
        $end = time();
        $c = ($end - $start) + ($end_point_m - $start_point_m);
        $total_time += $c;
    }
    printf("%s call(s)/s \n", (1/($total_time/$loop)));
}

function main(){
    $loop = 10000;
    $jsonString = loadJson();
    $jsonObj = json_decode($jsonString);
    decode($jsonString, $loop);
    encode($jsonObj, $loop);
}

main();

?>