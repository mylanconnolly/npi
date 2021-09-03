# frozen_string_literal: true

RSpec.describe NPI do
  it 'has a version number' do
    expect(NPI::VERSION).not_to be nil
  end

  describe '::valid?' do
    subject { described_class.valid?(value) }

    context 'when a non-numeric string is provided' do
      let(:value) { 'notannpiid' }

      it { is_expected.to eq(false) }
    end

    context 'when a string with an incorrect length is provided' do
      let(:value) { '13246' }

      it { is_expected.to eq(false) }
    end

    context 'when an integer with an incorrect length is provided' do
      let(:value) { 13_246 }

      it { is_expected.to eq(false) }
    end

    context 'when a valid NPI number string is used' do
      let(:value) { '1215290382' }

      it { is_expected.to eq(true) }
    end

    context 'when a valid NPI number integer is used' do
      let(:value) { 1_215_290_382 }

      it { is_expected.to eq(true) }
    end

    context 'when an invalid NPI number string is used' do
      let(:value) { '1215290383' }

      it { is_expected.to eq(false) }
    end

    context 'when an invalid NPI number integer is used' do
      let(:value) { 1_215_290_383 }

      it { is_expected.to eq(false) }
    end
  end
end
