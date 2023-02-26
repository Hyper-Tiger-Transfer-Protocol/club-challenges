package java_example;
import java.io.FileReader;

import org.json.simple.JSONArray;
import org.json.simple.JSONObject;
import org.json.simple.parser.*;

public class json_java
{
    public static void main(String[] args) throws Exception
    {
        Object obj = new JSONParser().parse(new FileReader("CryptoSeries/plaintext/plaintext_array.json"));
        JSONArray list = (JSONArray) obj;
        System.out.println(list.toString());

        obj = new JSONParser().parse(new FileReader("CryptoSeries/plaintext/plaintext_object.json"));
        JSONObject object = (JSONObject) obj;
        System.out.println(object.toString());
    }
}