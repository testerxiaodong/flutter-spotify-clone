import 'package:flutter/material.dart';

class CustomField extends StatelessWidget {
  final String hintText;
  final TextEditingController? controller;
  final bool isObscureText;
  final bool readonly;
  final VoidCallback? onTap;
  const CustomField({
    super.key,
    required this.hintText,
    required this.controller,
    this.isObscureText = false,
    this.readonly = false,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      onTap: onTap,
      controller: controller,
      readOnly: readonly,
      decoration: InputDecoration(hintText: hintText),
      obscureText: isObscureText,
      validator: (value) {
        if (value!.trim().isEmpty) {
          return "$hintText is missing";
        }
        return null;
      },
    );
  }
}
