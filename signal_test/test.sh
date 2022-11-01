cc capture.c -o capture -Wall -Wextra
./capture & 
ps aux | grep ./capture | awk {'print $2'} | xargs -I '{}' kill $* '{}'
rm capture

./capture.py & 
ps aux | grep ./capture.py | awk {'print $2'} | xargs -I '{}' kill $* '{}'
