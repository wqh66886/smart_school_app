import 'package:flutter/material.dart';
import 'package:smart_school_mobile/core/configs/theme/app_theme.dart';
import 'package:smart_school_mobile/pages/auth/sign_in_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: AppTheme.appTheme,
      home: const SignInPage(),
    );
  }
}