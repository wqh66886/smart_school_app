import 'package:flutter/material.dart';
import 'package:smart_school_mobile/pages/auth/sign_in_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  _border([Color color = Colors.grey]) => OutlineInputBorder(
    borderRadius: BorderRadius.circular(10),
    borderSide: BorderSide(
      color: Colors.grey,
      width: 3
    )
  );
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'smart school',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
        inputDecorationTheme: InputDecorationTheme(
          contentPadding: EdgeInsets.all(27),
          border: _border(),
          focusedBorder: _border(),
          errorBorder: _border(Colors.redAccent),
        ),
      ),
      home: const SignInPage(),
    );
  }
}