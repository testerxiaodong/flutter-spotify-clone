import 'package:client/core/theme/app_pallete.dart';
import 'package:client/core/widgets/custom_field.dart';
import 'package:dotted_border/dotted_border.dart';
import 'package:flex_color_picker/flex_color_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class UploadSongPage extends ConsumerStatefulWidget {
  const UploadSongPage({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _UploadSongPageState();
}

class _UploadSongPageState extends ConsumerState<UploadSongPage> {
  final songNameController = TextEditingController();
  final artistController = TextEditingController();
  Color selectedColor = Pallete.cardColor;

  @override
  void dispose() {
    super.dispose();
    songNameController.dispose();
    artistController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Upload Song'),
        actions: [IconButton(onPressed: () {}, icon: Icon(Icons.check))],
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            children: [
              DottedBorder(
                color: Pallete.borderColor,
                radius: const Radius.circular(10),
                borderType: BorderType.RRect,
                strokeCap: StrokeCap.round,
                dashPattern: const [10, 4],
                child: SizedBox(
                  height: 159,
                  width: double.infinity,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(Icons.folder_open, size: 40),
                      const SizedBox(height: 15),
                      Text('Select the thumnail for your song'),
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 40),
              CustomField(
                hintText: 'pick song',
                controller: null,
                readonly: true,
                onTap: () {},
              ),
              const SizedBox(height: 20),
              CustomField(hintText: 'Artist', controller: artistController),
              const SizedBox(height: 20),
              CustomField(hintText: 'Artist', controller: songNameController),
              const SizedBox(height: 20),
              ColorPicker(
                pickersEnabled: {ColorPickerType.wheel: true},
                color: selectedColor,
                onColorChanged: (Color color) {
                  setState(() {
                    selectedColor = color;
                  });
                },
                heading: Text('Select Color'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
