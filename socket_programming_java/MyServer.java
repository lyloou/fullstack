//: MyServer.java
import java.io.*;
import java.net.*;

public class MyServer{
  public static void main(String[]args) {
    try {
      ServerSocket ss = new ServerSocket(6666);
      Socket s = ss.accept();

      DataInputStream dis = new DataInputStream(s.getInputStream());
      DataOutputStream dos = new DataOutputStream(s.getOutputStream());

      BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

      String str = "";
      while (!str.equals("stop")) {

        str = (String)dis.readUTF();
        System.out.println("[client] " + str);

        str = br.readLine();
        System.out.println("[server] " + str);
        dos.writeUTF(str);
        dos.flush();

      }

      s.close();
      ss.close();

    } catch (Exception e){
      System.out.println(e);
    }
  }
}
