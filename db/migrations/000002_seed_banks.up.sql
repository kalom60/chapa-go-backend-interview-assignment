INSERT INTO bank (
  id, bank_id, slug, swift, name, acct_length, country_id,
  is_mobilemoney, is_active, is_rtgs, active, is_24hrs, currency, created_at, updated_at
) VALUES
  (gen_random_uuid(), 130, 'abay_bank', 'ABAYETAA', 'Abay Bank', 16, 1, 0, 1, 1, 1, 0, 'ETB', '2023-01-24T04:28:30Z', '2024-08-03T08:10:24Z'),
  (gen_random_uuid(), 772, 'addis_int_bank', 'ABSCETAA', 'Addis International Bank', 15, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 207, 'ahadu_bank', 'AHUUETAA', 'Ahadu Bank', 10, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 656, 'awash_bank', 'AWINETAA', 'Awash Bank', 14, 1, 0, 1, 0, 1, 0, 'ETB', '2022-03-17T04:21:30Z', '2024-08-02T20:08:46Z'),
  (gen_random_uuid(), 347, 'boa_bank', 'ABYSETAA', 'Bank of Abyssinia', 8, 1, 0, 1, 0, 1, 0, 'ETB', '2022-07-04T21:33:57Z', '2024-08-02T20:08:45Z'),
  (gen_random_uuid(), 571, 'berhan_bank', 'BERHETAA', 'Berhan Bank', 13, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 128, 'cbebirr', 'CBETETAA', 'CBEBirr', 10, 1, 1, 1, 0, 1, 1, 'ETB', '2024-01-24T14:41:12Z', '2024-08-12T20:16:07Z'),
  (gen_random_uuid(), 946, 'cbe_bank', 'CBETETAA', 'Commercial Bank of Ethiopia (CBE)', 13, 1, 0, 1, 0, 1, 1, 'ETB', '2022-03-17T04:21:18Z', '2024-08-03T05:56:23Z'),
  (gen_random_uuid(), 893, 'ebirr', 'CBORETA', 'Coopay-Ebirr', 10, 1, 1, 1, 0, 1, 1, 'ETB', '2023-08-15T08:00:11Z', '2024-08-10T14:30:16Z'),
  (gen_random_uuid(), 880, 'dashen_bank', 'DASHETAA', 'Dashen Bank', 13, 1, 0, 1, 0, 1, 0, 'ETB', '2022-11-15T03:17:43Z', '2024-08-02T20:08:46Z'),
  (gen_random_uuid(), 301, 'global_bank', 'DEGAETAA', 'Global Bank Ethiopia', 13, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 534, 'hibret_bank', 'UNTDETAA', 'Hibret Bank', 16, 1, 0, 1, 0, 1, 0, 'ETB', '2023-01-06T03:18:43Z', '2024-08-02T20:08:46Z'),
  (gen_random_uuid(), 315, 'anbesa_bank', 'LIBSETAA', 'Lion International Bank', 9, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 266, 'mpesa', 'MPESA', 'M-Pesa', 10, 1, 1, 1, 0, 1, 1, 'ETB', '2024-01-18T14:41:12Z', '2024-08-02T20:08:57Z'),
  (gen_random_uuid(), 979, 'nib_bank', 'NIBIETAA', 'Nib International Bank', 13, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 423, 'oromia_bank', 'ORIRETAA', 'Oromia International Bank', 12, 1, 0, 1, 1, 1, 1, 'ETB', '2024-08-12T04:21:18Z', '2024-08-12T04:21:18Z'),
  (gen_random_uuid(), 855, 'telebirr', 'TELEBIRR', 'telebirr', 10, 1, 1, 1, 0, 1, 1, 'ETB', '2022-12-12T14:41:12Z', '2024-08-02T20:08:57Z'),
  (gen_random_uuid(), 472, 'wegagen_bank', 'WEGAETAA', 'Wegagen Bank', 13, 1, 0, 1, 1, 1, 0, 'ETB', '2022-11-15T03:16:40Z', '2024-08-12T20:15:43Z'),
  (gen_random_uuid(), 687, 'zemen_bank', 'ZEMEETAA', 'Zemen Bank', 16, 1, 0, 1, 1, 1, 0, 'ETB', '2022-09-30T12:56:53Z', '2024-08-12T20:14:40Z');

