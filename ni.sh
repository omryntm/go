#!/bin/bash

# Fonksiyonlar
function secim_yap {
    echo "Hangi işlemi yapmak istersiniz? Lütfen bir sayı girin: "
    read secim
}

function devam_et {
    echo "Devam etmek için bir tuşa basın..."
    read devam
}

# Ana menü
while true
do
    echo "Netinternet Hızlı Komut Çalıştırma Sistemi"
    echo "------------------------------------------"
    echo "99. Öncelikle buradan sistem gerekliliklerini yükleyin !!!"
    echo "Lütfen yapmak istediğiniz işlemi seçin:"
    echo "1. KAYNAK KULLANIMI GÖRÜNTÜLE"
    echo "2. ANLIK SQL BAĞLANTILARINI GÖRÜNTÜLE"
    echo "3. NETWORK TRAFİK KULLANIMINI GÖRÜNTÜLE"
    echo "4. TÜM SERVİSLERİN DURUMLARINI GÖRÜNTÜLE"
    echo "5. APACHE EROR LOG DOSYANINDA GÜNCEL 10 SATIRI GÖRÜNTÜLE"
    echo "6. SUNUCUYA BAĞLI OLAN İP ADRESLERİNİ GÖRÜNTÜLE"
    echo "8. SUNUCUYA YAPILAN EN SON BAĞLANTIYI GÖSTER"
    echo "9. CPANEL LOG GÖRÜNTÜLE"
    echo "10. İŞLETİM SİSTEMİNİ ÖĞREN"
    echo "12. DİSK İSTATİSTİKLERİNİ GÖSTER"
    echo "13. ÇALIŞMA SÜRESİNİ GÖRÜNTÜLE"
    echo "15. CPANEL İŞLEMLERİ"
    echo "16. KURULUM İŞLEMLERİ"
    echo "0. ÇIKIŞ"

    secim_yap

    case $secim in
        1)
            htop
            devam_et
            ;;
        2)
            mysqladmin pro
            devam_et
            ;;
        3)
            ifconfig ens3 | grep "RX packets"
            devam_et
            ;;
        4)
            systemctl list-units --type=service --state=running
            devam_et
            ;;
        5)
            tail -n 10 /var/log/httpd/error_log
            devam_et
            ;;
        6)
            netstat -tan | grep ESTABLISHED | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -nr
            devam_et
            ;;
        8)
            last
            devam_et
            ;;
        
        99) yum install wget nano htop -y
            devam_et
            ;;
    9)
        echo "CPANEL LOG İŞLEMLERİ"
        echo "1 ACCESS LOG"
        echo "2 HİZMET DURUMU LOG"
        echo "3 DELİVERY MAİL LOG"
        echo "4 GELEN POSTA KUYRUĞU"
        echo "5 REJECT EDİLEN MESAJLAR"
        secim_yap
        case $secim in
        1)
           tail -n 20 /usr/local/cpanel/logs/access_log
                devam_et
                ;;
        2)
            tail -n 20 /var/log/chkservd.log
            devam_et
            ;;
        3)
            tail -n 20 /var/log/exim_mainlog
            devam_et
            ;;
        4)
            tail -n 20 /var/spool/exim/input/
            devam_et
            ;;
        5)
            tail -n 20 /var/log/exim_rejectlog
            devam_et
            ;;
          *)
          echo "Geçersiz seçim. Lütfen tekrar deneyin."
           devam_et
            ;;
            esac
            ;;
    10)
            uname -a
            devam_et
            ;;
        
        12)
            df -h
            devam_et
            ;;
        13)
            uptime
            devam_et
            ;;

        15)
            echo "CPANEL İŞLEMLERİ"
            echo "1 Aktif Backup İşlemini Görüntüle"
            echo "2 Tüm Hesaplar İçin backup işlemini Başlat"
            echo "3 Güncel Lisansı Kontrol Et"
            echo "4 Bilinen Servisleri Yeniden Başlat"
            echo "5 Servis Durumlarını Kontrol Et"
            secim_yap

            case $secim in
                1)
                    /usr/local/cpanel/scripts/cpbackup --list
                    devam_et
                    ;;
                2)
                    /usr/local/cpanel/bin/backup --force
                    devam_et
                    ;;
                3)
                    /usr/local/cpanel/cpkeyclt
                    devam_et
                    ;;
                4)
                    /scripts/restartsrv_httpd
                    /scripts/restartsrv_cpsrvd
                    /scripts/restartsrv_named
                    /scripts/restartsrv_mysqld
                    devam_et
                    ;;
                5)
                    service httpd status
                    devam_et
                    ;;
                *)
                    echo "Geçersiz seçim. Lütfen tekrar deneyin."
                    devam_et
                    ;;
            esac
            ;;
        16)
            echo "KURULUM İŞLEMLERİ"
            echo "!!! KURULAN PANELLER KALDIRILMADAN YENİ PANEL KURULAMAZ !!!"
            #echo "1 Yeni bir kullanıcı oluştur"
            #echo "2 Yeni bir dizin oluştur"
            echo "1 Cpanel kurulumu yap"
            echo "2 Directadmin kurulumu yap"
            echo "3 Plesk kurulumu yap"
            echo "4 CWP kurulumu yap"
            echo "5 Cyberpanel kurulumu yap"
            echo "6 Fastpanel kurulumu yap"
            secim_yap

            case $secim in
                #1)
                    #echo "Yeni kullanıcı adı: "
                    #read username
                    #adduser $username
                    #passwd $username
                    #devam_et
                    #;;
                #2)
                    #echo "Yeni dizin adı: "
                    #read dirname
                    #mkdir $dirname
                    #devam_et
                    #;;
                
                1) cd /home && curl -o latest -L https://securedownloads.cpanel.net/latest && sh latest
                    devam_et
                    ;;
                
                2) 
                    echo "Lisans_anahtarı: "
                    read Lisans_anahtarı
                    cd /home && wget https://download.directadmin.com/setup.sh 'Lisans_anahtarı'
                    devam_et
                    ;;

                3) cd /home && wget http://autoinstall.plesk.com/plesk-installer && chmod +x plesk-installer && ./plesk-installer
                    devam_et
                    ;;
                
                4) cd /home && wget http://centos-webpanel.com/cwp-el7-latest && chmod +x cwp-el7-latest && ./cwp-el7-latest
                    devam_et
                    ;;
                
                5) cd /home && wget -O - http://cyberpanel.sh | sh
                    devam_et
                    ;;
                
                6) cd /home && wget wget http://repo.fastpanel.direct/install_fastpanel.sh -O - | bash -
                    devam_et
                    ;;


                *)
                    echo "Geçersiz seçim. Lütfen tekrar deneyin."
                    devam_et
                    ;;
            esac
            ;;
        0)
            echo "Programdan çıkılıyor..."
            exit 0
            ;;
        *)
            echo "Geçersiz seçim. Lütfen tekrar deneyin."
            devam_et
            ;;
    esac
done