#!/bin/bash
nama=$1
tanggal="$(date)"

root_folder="$nama_$tanggal"

mkdir -p "${root_folder}/about_me/personal"
akun_fb=$2
echo "https://www.facebook.com/${akun_fb}" >> "${root_folder}/about_me/personal/facebook.txt"

mkdir -p "${root_folder}/about_me/professional"
akun_lnkd=$2
echo "https://www.linkedin.com/in/${akun_lnkd}" >> "${root_folder}/about_me/professional/linkedin.txt"

mkdir -p "${root_folder}/my_friends"
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt >> "${root_folder}/my_friends/list_of_my_friends.txt"

mkdir -p "${root_folder}/my_system_info"
my_username=$(hostname)
detail_info=$(uname -a)
echo "My username: ${my_username}" >> "${root_folder}/my_system_info/about_this_laptop.txt"

echo "With host: ${detail_info}" >> "${root_folder}/my_system_info/about_this_laptop.txt"

google_ping=$(ping -c 3 google.com)
echo "Connection to google:" >> "${root_folder}/my_system_info/internet_connection.txt"
echo "${google_ping}" >> "${root_folder}/my_system_info/internet_connection.txt"