package info.sqldev.sign;

import java.security.MessageDigest;
import java.util.*;

public class Signer {
    private static String encodeHex(byte[] data) {
        int l = data.length;
        byte[] out = new byte[l << 1];
        int i = 0;
        byte[] toDigits = new byte[]{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
        for (int var5 = 0; i < l; i++) {
            out[var5] = toDigits[(240 & data[i]) >> 4];
            var5++;
            out[var5] = toDigits[15 & data[i]];
            var5++;
        }
        return new String(out);
    }

    public static String getSignature(Map<String, String> params, String appSecret) throws Exception {
        TreeMap<String, String> sortedParams = new TreeMap<String, String>();
        for (Map.Entry<String, String> entry : params.entrySet()) {
            if (entry.getValue() != null && entry.getValue().length() > 0) {
                sortedParams.put(entry.getKey(), entry.getValue());
            }
        }
        StringBuilder baseString = new StringBuilder();
        for (Map.Entry<String, String> entry : sortedParams.entrySet()) {
            if (entry.getKey() != null && entry.getKey() != "sign" && entry.getKey() != "key" && entry.getValue() != null) {
                baseString.append(entry.getKey());
                baseString.append("=");
                baseString.append(entry.getValue());
                baseString.append("&");
            }
        }
        if (baseString.length() > 0) {
            baseString.deleteCharAt(baseString.length() - 1);
            baseString.append(appSecret);
        }

        return encodeHex(MessageDigest.getInstance("MD5").digest(baseString.toString().getBytes("UTF-8")));
    }
}