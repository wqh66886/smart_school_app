import 'package:flutter/material.dart';
import 'package:smart_school_mobile/widgets/auth_gradient_button.dart';
import 'package:smart_school_mobile/widgets/text_editing.dart';

class SignInPage extends StatefulWidget {
  const SignInPage({super.key});

  @override
  State<SignInPage> createState() => _SignInPageState();
}

class _SignInPageState extends State<SignInPage> {
  final textController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(27.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Image.asset("assets/images/school.png"),
            SizedBox(
              height: 20,
            ),
            Text(
              "智慧校园",
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            SizedBox(
              height: 15,
            ),
            Text("开启校园美好生活",
                style: TextStyle(fontSize: 14, fontWeight: FontWeight.w400)),
            SizedBox(
              height: 20,
            ),
            TextEditing(
                hintText: "邮箱/手机",
                prefixIcon: Icons.people_alt_outlined,
                controller: textController),
            SizedBox(
              height: 15,
            ),
            TextEditing(
              hintText: "密码",
              prefixIcon: Icons.lock_clock_outlined,
              controller: passwordController,
              isObscureText: true,
            ),
            SizedBox(
              height: 15,
            ),
            AuthGradientButton(
              onPressed: () {},
              buttonText: "登录",
            ),
          ],
        ),
      ),
    );
  }
}
