<?php

echo json_encode(get_loaded_extensions(), JSON_PRETTY_PRINT) . PHP_EOL;;

echo \SConcur\Extension\ping('hello') . PHP_EOL;