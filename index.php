<?php

// wczytanie pliku csv oraz jego konwersja do array

$csv = array_map('str_getcsv', file('php_internship_data.csv')); 

$names = array_column($csv, 0);

$iterr = array_count_values($names);
arsort($iterr);
$top10 = array_slice($iterr, 0 , 10);



foreach ($top10 as $name => $count) {
    
   $name = ucfirst(mb_strtolower($name, "UTF-8"));

  echo nl2br("$name = $count \n ");
}




?>