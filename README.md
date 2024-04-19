# WiFi Cihaz Listesi API

Bu basit Go programı, bir ağdaki WiFi cihazlarının listesini sunan bir HTTP API sağlar.

## Kullanım

- Program, `/wifi/devices` endpoint'ına yapılan GET isteklerini dinler.
- Endpoint'a yapılan istekler, ağdaki cihazların MAC adreslerini ve IP adreslerini döndürür.
- Cihaz listesi, ARP tablosundan alınır.
- Program, 8080 portunda bir HTTP sunucusu olarak çalışır.

## Dikkat

- Programın çalışabilmesi için `arp` komutu gereklidir.
- Program, ARP tablosunu kullanarak cihazları tespit eder.
- Programı çalıştırmadan önce, gerekli izinleri ve bağımlılıkları sağladığınızdan emin olun.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
