#!/usr/bin/perl -n

if (m/^(\w+) (\w+) (\(.*?\))/) {
	$n=$j=$1;
	$t=$2;
	$o=$d=$3;

	$n=~s/_/ /g;
	$n=~s/\b(\w)/\U$1/g;
	$n=~s/ //g;

	if ($t eq "map") {
		$t = "device";
	} elsif ($t eq "device_class") {
		$t = "DeviceClass";
	}
	$t = ucfirst($t);

	$d=~s/^.*?default:\s*(\w+).*/ default:"$1"/;
	if ($d eq "(optional)") {
		$d = "";
	}
	if ($o eq "(optional)") {
		$o = "";
	}

	$p = sprintf("$n $t \x60json:\"$j,omitempty\"$d\x60\n");
	# printf("$n $t \x60json:\"$j,omitempty\"$d\x60\t// $o\n// $j $t $o\n");

} elsif (m/^(\w+) (\w+) (REQUIRED)/) {
	$n=$j=$1;
	$t=$2;
	$o=$3;

	$n=~s/_/ /g;
	$n=~s/\b(\w)/\U$1/g;
	$n=~s/ //g;

	$t = ucfirst($t);

	$p = sprintf("$n $t \x60json:\"$j,omitempty\" required:\"true\"\x60\n");
	# printf("$n $t \x60json:\"$j,omitempty\" required:\"true\"\x60\n// $j $t $o\n");

} elsif (/^\s+$/) {
	printf("\n");

} else {
	printf("// $_$p");
	$p = "";
}

