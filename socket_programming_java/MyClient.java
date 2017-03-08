//: MyClient.java
import java.io.*;
import java.net.*;

public class MyClient {
  public static void main (String []args) {
    try {
      Socket s = new Socket("localhost", 6666);
      DataOutputStream dos = new DataOutputStream(s.getOutputStream());
      DataInputStream dis = new DataInputStream(s.getInputStream());

      BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
      String str = "";

      while(!str.equals("stop")) {
        str = br.readLine();
        dos.writeUTF(str);
        dos.flush();
        System.out.println("[client] " + str);

        System.out.println("[server] " + dis.readUTF());
      }

      dos.close();
      s.close();
    } catch (Exception e) {
      System.out.println(e);
    }

  }
}
