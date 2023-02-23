package java_example;
import java.io.FileReader;

import org.json.simple.JSONArray;
import org.json.simple.parser.*;

public class json_java
{
    public static void main(String[] args) throws Exception
    {
        // parsing file "JSONExample.json"
        Object obj = new JSONParser().parse(new FileReader("CryptoSeries/plaintext/plaintext.json"));

        // typecasting obj to JSONObject
        JSONArray jo = (JSONArray) obj;

        System.out.println(jo.toString());
    }
}