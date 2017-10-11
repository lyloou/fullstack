import java.net.*;

public class DSender {
  public static void main(String[] args) {
    try {
      DatagramSocket ds = new DatagramSocket();
      String str = "Welcome java";
      InetAddress ia = InetAddress.getByName("127.0.0.1");

      DatagramPacket dp = new DatagramPacket(str.getBytes(), str.length(), ia, 3000);
      ds.send(dp);
      ds.close();
    } catch (Exception e) {
      System.out.println(e);
    }
  }
}
