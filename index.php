<?php

$file = 'php_internship_data.csv';

// sprawdzenie obecności zadeklarowanego pliku oraz wczytanie i konwersja do kolekcji
if (!file_exists($file)) {
  echo "brak pliku ", $file;
  exit;
}
$csv = array_map('str_getcsv', file($file));

// wydzielenie listy imion z kolekcji, obliczenie ilości wystąpień tego samego imienia oraz sortowanie od największej
$names = array_column($csv, 0);
$iterrations = array_count_values($names);
arsort($iterrations);

// wyciecie z kolekcji 10 pierwszych wyników
$top10 = array_slice($iterrations, 0, 10);

// wyświetlenie wyniku końcowego poprzez rozbicie kolekcji na imię i liczbę wystąpień
echo "<h1>";
echo "Top 10 występujących imion:";
echo "</h1>";
foreach ($top10 as $name => $count) {

  $name = ucfirst(mb_strtolower($name, "UTF-8"));

  echo nl2br("$name = $count \n");
}

?>