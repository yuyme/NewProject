package learn.DataType;

import java.util.Scanner;

public class arrDemo {
    public static void main(String[] args) {
        int[] arr = {0,1,2,2,3,0,4,2};


        // isExce(arr);

        int[] arrs = numArr(arr, 2);
        for (int i : arrs) {
            System.out.print(i + " ");
        }
    }

    private static void isExce(int[] arr) {
        Scanner sc = new Scanner(System.in);
        System.out.println("请输入任意一个两位整数：");
        int num = sc.nextInt();

        boolean flag = false;
        int index = 0;

        for (int i = 0; i < arr.length; i++) {
            if (num == arr[i]) {
                flag = true;
                index = i;
                break;
            }
        }
        if (flag) System.out.println(index);
        else System.out.println("该数据不存在");
    }

    private static int[] numArr(int[] arr, int val) {
        int count = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == val){
                arr[i] = -1;
                count++;
            }
        }

        int[] newArr = new int[arr.length - count];
        int index = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] != -1){
                newArr[index] = arr[i];
                index++;
            }
        }

        return newArr;
    }
}
