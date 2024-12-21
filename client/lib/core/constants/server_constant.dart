import 'dart:io';

class ServerConstant {
  static String serverUrl =
      Platform.isAndroid ? 'http://10.0.2.2:8080' : 'http://192.168.5.241:8080';
}
