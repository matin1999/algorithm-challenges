<?php

$strs = ["eat", "tea", "tan", "ate", "nat", "bat"];

$groups = [];

foreach ($strs as $str) {
    $chars = str_split($str);
    sort($chars);
    $key = implode("", $chars);

    if (!isset($groups[$key])) {
        $groups[$key] = [];
    }

    $groups[$key][] = $str;
}

foreach ($groups as $group) {
    print_r($group);
    echo "\n";
}